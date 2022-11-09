package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	app.Get("/", hello)
	app.Get("/panic", withPanic)

	log.Fatal(app.Listen(":3000"))
}

func withPanic(c *fiber.Ctx) error {
	panic("Some error")
}

func hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
