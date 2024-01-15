package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var db *sql.DB;

func main() {

    if err := godotenv.Load(); err != nil {
        os.Stderr.WriteString(err.Error());
    }
    
    db, err := configureDatabase();
    if err != nil {
        os.Stderr.WriteString(err.Error());
    }

    defer db.Close();

    if err := db.Ping(); err != nil {
        os.Stderr.WriteString(err.Error());
    }


    app := fiber.New();

    app.Post("/register", func(c *fiber.Ctx) error {
        return register(c, db);
    });

    app.Get("/verify/:user_id", func(c *fiber.Ctx) error {
        return verify(c, db);
    })

    app.Post("/login", func(c *fiber.Ctx) error {
        return login(c, db);
    });

    app.Post("/image", func(c *fiber.Ctx) error {
        return saveImage(c);
    });


    app.Listen(":" + os.Getenv("SERVER_PORT"));
}

