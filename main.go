package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/trioxtron/htmx-tic-tac-toe/gameengine"
)

func setupRoutes(app *fiber.App) {
    game := &gameengine.Game{}
	app.Static("/src/tailwind.css", "./src/tailwind.css")
    app.Get("/", func(c *fiber.Ctx) error {
        board := game.GetBoard()
        return c.Render("index", fiber.Map{
            "f1": board[0][0],
            "f2": board[0][1],
            "f3": board[0][2],
            "f4": board[1][0],
            "f5": board[1][1],
            "f6": board[1][2],
            "f7": board[2][0],
            "f8": board[2][1],
            "f9": board[2][2],
        })
    })
    app.Post("/setfield", func(c *fiber.Ctx) error {
        field := struct {
            Y int `json:"y"`
            X int `json:"x"`
        }{}
        if err := c.BodyParser(&field); err != nil {
            fmt.Println("error = ", err)
            return c.SendStatus(200)
        }
        fmt.Println(field)
        return c.SendString(game.SetField(field.Y, field.X))
	})
    app.Post("/clearfield", func(c *fiber.Ctx) error {
        game.ClearField()
        return c.SendStatus(200)
    })
}

func main() {
    engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))

	setupRoutes(app)
	app.Listen(":3000")
}
