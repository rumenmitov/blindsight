package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
)

var db *sql.DB;

func main() {
    
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

    app.Post("/login", func(c *fiber.Ctx) error {
        return login(c, db);
    });

    app.Post("/image", func(c *fiber.Ctx) error {
        return saveImage(c);
    });



    app.Listen(":3000");
}

