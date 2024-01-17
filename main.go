package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var db *sql.DB;

func main() {

    if err := godotenv.Load(); err != nil {
        Log(err.Error());
    }
    
    db, err := configureDatabase();
    if err != nil {
        Log(err.Error());
    }

    defer db.Close();

    if err := db.Ping(); err != nil {
        Log(err.Error());
    }


    app := fiber.New();

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("pong");
    });

    app.Get("/users", func(c *fiber.Ctx) error {
        users, err := users(db);
        if err != nil {
            Log(err.Error());
        }

        return c.SendString(string(users));
    });

    app.Post("/register", func(c *fiber.Ctx) error {
        err := register(c, db);
        if err != nil {
            Log(err.Error());
        }

        return err;
    });

    app.Post("/verify", func(c *fiber.Ctx) error {
        userJSON, err := verify(c, db);
        if AuthError(err) != AuthError(Ok) {
            Log(fmt.Sprintf("Verification error: %d", AuthError(err)));

            // NOTE: 201 is the status for a CREATED response. 201+ is an error.
            c.Status(int(AuthError(err)))
            return nil;
        }

        c.Status(201);
        return c.SendString(string(userJSON));
    })

    app.Post("/login", func(c *fiber.Ctx) error {
        err := login(c, db);
        if err != nil {
            Log(err.Error());
        }

        return err;
    });

    app.Post("/image", func(c *fiber.Ctx) error {
        return saveImage(c);
    });


    app.Listen(":" + os.Getenv("SERVER_PORT"));
}

