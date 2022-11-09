package demo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func NewHello(app *fiber.App, client *mongo.Client) {
	app.Get("/", helloHandler(client))
}

var tracer = otel.Tracer("demo-api")

func helloHandler(client *mongo.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, span := tracer.Start(c.Context(), "helloHandler", oteltrace.WithAttributes(attribute.String("layer", "handler")))
		defer span.End()

		return c.JSON(fiber.Map{
			"message": "Hello API",
		})
	}
}
