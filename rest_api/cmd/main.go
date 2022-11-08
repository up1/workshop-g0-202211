package main

import (
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
	client := demo.NewMongoClient(*url)

	app := demo.NewRouter()
	demo.NewHello(app, client)

	c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        go func() {
                <-c
                fmt.Println("Gracefully shutting down...")
                _ = app.Shutdown()
        }()

	log.Fatal(app.Listen(":3000"))
}
