package main

import (
	"encoding/base64"
	"errors"
	"os"
	"github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

func register(c *fiber.Ctx) error {

    db, err := configureDatabase();
    if err != nil {
        return err;
    }

    defer db.Close();

    hashed_pass, err := 
        bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost);

    if err != nil {
        return err;
    }

    register_user := `insert into "Users"("fname", "lname", "email", "username", "password") 
        values($2, $3, $4, $5, $6)`;

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
