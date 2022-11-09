package main

import (
	"day03"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

var server = flag.String("server", ":3000", "Host") 

func main() {
	flag.Parse()
	p := os.Getenv("PORT")
	fmt.Println(p, *server)

	app := day03.NewFiberRouter()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

   // init dependency
	r := day03.HelloRepository{}
	s := day03.NewService(r)

	day03.NewHelloRouter(app, s)
	app.Get("/panic", withPanic)

	log.Fatal(app.Listen(*server))
}

func withPanic(c *fiber.Ctx) error {
	panic("Some error")
}
