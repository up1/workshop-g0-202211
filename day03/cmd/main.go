package main

import (
	"day03"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := day03.NewFiberRouter()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

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
