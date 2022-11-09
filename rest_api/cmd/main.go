package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"demo"
)

var url = flag.String("url", "mongodb://user:pass@128.199.205.113:27017", "Host")

func main() {
	// Connect to mongodb
	flag.Parse()
	client, _ := demo.NewMongoClient(*url)

	// Logging
	demo.InitializeLogger()

	// Tracer
	tp, err := demo.InitTracer("http://128.199.205.113:9411/api/v2/spans", "demo-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		tp.Shutdown(context.Background())
	}()

	app := demo.NewRouter()
	demo.NewHello(app, client)
	demo.NewHealthCheck(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	log.Fatal(app.Listen(":3000"))
}
