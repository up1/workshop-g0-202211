package main

import (
	"context"
	"demo/other"
	"log"

	"demo"
)

func main() {
	// Tracer
	tp, err := demo.InitTracer("http://128.199.205.113:9411/api/v2/spans", "other-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		tp.Shutdown(context.Background())
	}()

	app := demo.NewRouterOther()
	other.NewOther(app)

	log.Fatal(app.Listen(":4000"))
}
