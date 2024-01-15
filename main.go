package main

import (
	"github.com/gofiber/fiber/v2"
)


func setupRoutes (app *fiber.App) {
    app.Static("/src/tailwind.css", "./src/tailwind.css")
    app.Static("/", "./views/index.html")
    app.Post("/setfield", func(c *fiber.Ctx) error {
        return c.SendString("Nice turn!")
    })
}

func main () {
    app := fiber.New()

    setupRoutes(app)
    app.Listen(":3000")
}
