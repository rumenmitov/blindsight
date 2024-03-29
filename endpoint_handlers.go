package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"math/rand"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func users(db *sql.DB) ([]byte, error) {
    get_users := `SELECT "username", "email" FROM "Users" WHERE "verified" = true`;

    results, err := db.Query(get_users);
    if err != nil {
        return nil, err;
    }

    defer results.Close();

    var users []User;

    for results.Next() {
        var user User;

        if err := results.Scan(&user.Username, &user.Email); err != nil {
            return nil, err;
        }

        users = append(users, user);
    }

    return json.Marshal(users);
}

func register(c *fiber.Ctx, db *sql.DB) error {
    create_db_if_none_exists := `CREATE TABLE IF NOT EXISTS "Users"(
        "id" SERIAL PRIMARY KEY,
        "fname" VARCHAR(100) NOT NULL,
        "lname" VARCHAR(100) NOT NULL,
        "email" VARCHAR(100) UNIQUE NOT NULL,
        "username" VARCHAR(100) UNIQUE NOT NULL,
        "password" VARCHAR(100) NOT NULL,
        "verification_code" INT,
        "verified" BOOLEAN
    )`;

    _, err := db.Exec(create_db_if_none_exists);
    if err != nil {
        return err;
    }

    hashed_pass, err := 
        bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost);

    if err != nil {
        return err;
    }

    // NOTE: For uniformity, all verification codes should be 6-digit
    verification_code := rand.Intn(899999) + 100000;

    register_user := `INSERT INTO "Users"("fname", "lname", "email", "username", "password", "verification_code", "verified") 
        VALUES($1, $2, $3, $4, $5, $6, $7)`;

    _, err = db.Exec(register_user, 
    c.FormValue("fname"),
    c.FormValue("lname"),
    c.FormValue("email"),
    c.FormValue("username"),
    string(hashed_pass),
    verification_code,
    false );

    if err != nil {
        return err;
    }

    verification_mes := gomail.NewMessage();
    verification_mes.SetHeader("From", os.Getenv("EMAIL_USER"));
    verification_mes.SetHeader("To", c.FormValue("email"));
    verification_mes.SetHeader("Subject", "Verify your Account");
    verification_mes.SetBody("text/html", ` 
        <h1>Welcome to BlindSight, ` + c.FormValue("fname") + ` ` + c.FormValue("lname") +`!</h1>
        <p>To complete your registration, please enter the following code in the BlindSight app:</p>
        <p><b>` + fmt.Sprint(verification_code) + `</b></p>
        <p>If you do not remember creating an account with BlindSight, please ignore this email.</p>
        <br>
        <p>Sincerely,</p>
        <p>The BlindSight Team</p>
    `);

    email_dialer := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"));

    if err := email_dialer.DialAndSend(verification_mes); err != nil {
        return err;
    }

    return nil;
    
}

func verify(c *fiber.Ctx, db *sql.DB) ([]byte, AuthError) {

    verify_user := `UPDATE "Users" SET "verified" = true, "verification_code" = NULL WHERE "verification_code" = $1
                    RETURNING "fname", "lname", "email", "username", "password"`;

    verification_code, err := strconv.Atoi(c.FormValue("verification_code"));
    if err != nil {
        return nil, AuthError(NotANumberError);
    }

    results, err := db.Query(verify_user, verification_code);
    if err != nil {
        return nil, AuthError(WrongCredentialsError);
    }

    for results.Next() {
        var user User;

        results.Scan(&user.Fname, &user.Lname, &user.Email, &user.Username, &user.Password);

        userJSON, err := json.Marshal(user);
        if err != nil {
            Log(err.Error())
            return nil, AuthError(UnkownError);
        }

        return userJSON, AuthError(Ok);

    }

    return nil, AuthError(UnkownError);
}

func login(c *fiber.Ctx, db *sql.DB) ([]byte, AuthError) {

    results, err := db.Query(`SELECT DISTINCT 
                              "fname",
                              "lname",
                              "email",
                              "username",
                              "password"
                              FROM "Users"
                              WHERE ("username" = '` + c.FormValue("username") +
                              `' OR "email" = '` + c.FormValue("username") + `') AND 
                              "verified" = true`);

    if err != nil {
        Log(err.Error());
        return nil, AuthError(WrongCredentialsError);
    }

    for results.Next() {
        var user User;

        if err := results.Scan(&user.Fname, &user.Lname, &user.Email, 
            &user.Username, &user.Password); err != nil {

                Log(err.Error());
                return nil, AuthError(UnkownError);
        }

        password_match := bcrypt.CompareHashAndPassword(
            []byte(user.Password), 
            []byte(c.FormValue("password")));

            // NOTE: Checks if no errors ocurred i.e. passwords match
        if password_match == nil {
            user_data, err := json.Marshal(user);
            if err != nil { 
                Log(err.Error());
                return nil, AuthError(UnkownError);
            }

            return user_data, AuthError(Ok);

        } else {
            return nil, AuthError(WrongCredentialsError);
        }

    }

    return nil, AuthError(UnkownError);
}


func saveImage(c *fiber.Ctx) ([]byte, error) {
    image := ImageFile {
        name: c.FormValue("name"),
        bytes: c.FormValue("image"),
    };

    dec, err := base64.StdEncoding.DecodeString(image.bytes);
    if err != nil {
        return nil, errors.New("Couldn't decode file!\n");
    }

    const input_dir = "/midas/input/";
    const output_dir = "/midas/output/";
    const model = "dpt_swin2_tiny_256";
    const model_path = "/midas/weights/dpt_swin2_tiny_256.pt";

    _, err = os.Stat(input_dir);
    if err != nil {
        if os.IsNotExist(err) {
            os.Mkdir(input_dir, os.ModePerm);
        } else {
            return nil, err;
        }
    }

    file, err := os.Create(input_dir + image.name + ".png");
    if err != nil {
        return nil, errors.New("Couldn't create file! \n");
    }

    defer file.Close();

    _, err = file.Write(dec);
    if err != nil {
        return nil, errors.New("Couldn't write to file!\n");
    }

    err = file.Sync();
    if err != nil {
        return nil, errors.New("Couldn't sync file to disk!\n");
    }

    time.Sleep(time.Millisecond * 1500);

    defer os.Remove(output_dir + image.name + "-" + model + ".png");

    instruction, err := exec.Command("depth_analyzer", output_dir + image.name + "-" + model + ".png").Output();
    if err != nil {
        return nil, err;
    }


    return instruction, nil;
}
