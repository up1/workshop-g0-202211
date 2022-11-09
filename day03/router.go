package day03

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiberRouter() *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	return app
}
