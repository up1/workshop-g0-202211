package day03

import "github.com/gofiber/fiber/v2"

func NewHelloRouter(app *fiber.App) {
	app.Get("/", Hello)
}

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
