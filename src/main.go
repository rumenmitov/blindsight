package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

    app := fiber.New();

    app.Post("/register", register);
    app.Post("/image", saveImage);



    app.Listen(":3000");
}

