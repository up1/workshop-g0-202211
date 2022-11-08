package demo

import "github.com/gofiber/fiber/v2"

func NewHello(app *fiber.App) {
	app.Get("/", helloHandler)
}

func helloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello API",
	})
}
