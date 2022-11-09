package day03

import "github.com/gofiber/fiber/v2"

func NewHelloRouter(app *fiber.App) {
	app.Get("/", hello)
}

func hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
