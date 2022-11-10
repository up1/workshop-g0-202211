package demo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func NewHello(app *fiber.App, client *mongo.Client) {
	app.Get("/", helloHandler(client))
}

var tracer = otel.Tracer("demo-api")

func helloHandler(client *mongo.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx, span := tracer.Start(c.UserContext(), "helloHandler", oteltrace.WithAttributes(attribute.String("layer", "handler")))
		defer span.End()

		Logger.Info("demo", zap.String("traceId", span.SpanContext().TraceID().String()))

		CallExternalApi(ctx)

		return c.JSON(fiber.Map{
			"message": "Hello API",
		})
	}
}

func CallExternalApi(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "CallExternalApi")
	defer span.End()

	var body []byte
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	tr := otel.Tracer("hello/gateway")
	err := func(ctx context.Context) error {
		ctx, span := tr.Start(ctx, "say hello", oteltrace.WithAttributes(semconv.PeerServiceKey.String("other-service")))
		defer span.End()

		ctx = httptrace.WithClientTrace(ctx, otelhttptrace.NewClientTrace(ctx))
		req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:4000", nil)

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		body, err = ioutil.ReadAll(res.Body)
		_ = res.Body.Close()

		return err
	}(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response Received: %s\n", body)
}
