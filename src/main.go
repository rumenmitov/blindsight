package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New();

    app.Post("/image", func(c *fiber.Ctx) error {

        return saveImage(c.FormValue("name"), c.FormValue("image"));

    });

    app.Listen(":3000");
}

