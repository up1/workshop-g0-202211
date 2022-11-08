package demo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHello(app *fiber.App, client *mongo.Client) {
	app.Get("/", helloHandler(client))
}

func helloHandler(client *mongo.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello API",
		})
	}
}
