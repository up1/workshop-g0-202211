package demo

import (
	"github.com/gofiber/fiber/v2"
)

func NewHealthCheck(app *fiber.App) {
	app.Get("/health", checkDependencies)
}

func checkDependencies(c *fiber.Ctx) error {
	// Check MongoDB
	client, err := NewMongoClient("mongodb://user:pass@128.199.205.113:27017")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"mongodb": "Not OK",
		})
	}
	defer Disconnect(client)

	return c.JSON(fiber.Map{
		"mongodb": "OK",
	})
}
