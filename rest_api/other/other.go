package other

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

func NewOther(app *fiber.App) {
	app.Get("/", otherHandler)
}

var tracer = otel.Tracer("other-api")

func otherHandler(c *fiber.Ctx) error {
	_, span := tracer.Start(c.UserContext(), "other-service")
	defer span.End()

	return c.JSON(fiber.Map{
		"message": "Call other api",
	})
}
