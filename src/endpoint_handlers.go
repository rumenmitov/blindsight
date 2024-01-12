package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
    "database/sql"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func register(c *fiber.Ctx, db *sql.DB) error {

    create_db_if_none_exist := `CREATE TABLE IF NOT EXISTS "Users"(
        "id" SERIAL PRIMARY KEY,
        "fname" VARCHAR(100) NOT NULL,
        "lname" VARCHAR(100) NOT NULL,
        "email" VARCHAR(100) UNIQUE NOT NULL,
        "username" VARCHAR(100) UNIQUE NOT NULL,
        "password" VARCHAR(100) NOT NULL
    )`;

    _, err := db.Exec(create_db_if_none_exist);
    if err != nil {
        return err;
    }

    hashed_pass, err := 
        bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost);

    if err != nil {
        return err;
    }

    register_user := `INSERT INTO "Users"("fname", "lname", "email", "username", "password") 
        VALUES($1, $2, $3, $4, $5)`;

    _, err = db.Exec(register_user, 
        c.FormValue("fname"),
        c.FormValue("lname"),
        c.FormValue("email"),
        c.FormValue("username"),
        string(hashed_pass) );

    if err != nil {
        return err;
    }

    return nil;
}

func login(c *fiber.Ctx, db *sql.DB) error {

    db, err := configureDatabase();
    if err != nil {
        return err;
    }

    defer db.Close();

    results, err := db.Query(`SELECT DISTINCT * FROM "Users"
                              WHERE username = ` + c.FormValue("username") +
                              `OR email = ` + c.FormValue("username"));

    if err != nil {
        return err;
    }

    for results.Next() {
        var user User;

        if err := results.Scan(&user.Fname, &user.Lname, &user.Email, 
            &user.Username, &user.Password); err != nil {

                return err;
        }

        password_match := bcrypt.CompareHashAndPassword(
            []byte(user.Password), 
            []byte(c.FormValue("password")));

        // checks if no errors ocurred i.e. passwords match
        if password_match == nil {
            user_data, err := json.Marshal(user);
            if err != nil { return err }

            return c.SendString(string(user_data));
        }

    }

    return nil;
}


func saveImage(c *fiber.Ctx) error {
    image := ImageFile {
        name: c.FormValue("name"),
        bytes: c.FormValue("image"),
    };

    dec, err := base64.StdEncoding.DecodeString(image.bytes);
    if err != nil {
        return errors.New("Couldn't decode file!\n");
    }

    _, err = os.Stat("images/");
    if err != nil {
        if os.IsNotExist(err) {
            os.Mkdir("images", os.ModePerm);
        } else {
            return err;
        }
    }

    file, err := os.Create("images/" + image.name + ".png");
    if err != nil {
        return errors.New("Couldn't create file!\n");
    }

    defer file.Close();

    _, err = file.Write(dec);
    if err != nil {
        return errors.New("Couldn't write to file!\n");
    }

    err = file.Sync();
    if err != nil {
        return errors.New("Couldn't sync file to disk!\n");
    }

    return nil;
}
