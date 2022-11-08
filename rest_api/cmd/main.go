package main

import (
	"log"

	"demo"
)

func main() {
	app := demo.NewRouter()
	demo.NewHello(app)
	log.Fatal(app.Listen(":3000"))
}
