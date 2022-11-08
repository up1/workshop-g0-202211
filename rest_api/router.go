package demo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter() *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	return app
}
