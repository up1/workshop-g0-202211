package main

import (
	"context"
	"day03"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var server = flag.String("server", ":3000", "Host")

func main() {
	flag.Parse()

	// Init tracer
	tp, err := day03.InitTracer("http://128.199.205.113:9411/api/v2/spans", "demo-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		tp.Shutdown(context.Background())
	}()

	day03.InitializeLogger()

	day03.Logger.Info("failed to fetch URL",
		zap.String("url", "xxxx.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	client := day03.NewMongoClient("mongodb://user:pass@128.199.205.113:27017")

	p := os.Getenv("PORT")
	fmt.Println(p, *server)

	app := day03.NewFiberRouter()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		day03.DisconnectMongoDB(client)
		_ = app.Shutdown()
	}()

	// init dependency
	r := day03.HelloRepository{
		Client: client,
	}
	s := day03.NewService(r)

	day03.NewHelloRouter(app, s)
	app.Get("/panic", withPanic)

	log.Fatal(app.Listen(*server))
}

func withPanic(c *fiber.Ctx) error {
	panic("Some error")
}
