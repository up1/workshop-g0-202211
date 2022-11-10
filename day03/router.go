package day03

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiberRouter() *fiber.App {
	app := fiber.New()

	app.Use(otelfiber.Middleware("demo-service"))

	prometheus := fiberprometheus.New("demo-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Use(recover.New())
	return app
}
