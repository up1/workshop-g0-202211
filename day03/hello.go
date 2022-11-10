package day03

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func NewHelloRouter(app *fiber.App, s HelloSevice) {
	app.Get("/", Hello(s))
}

var tracer = otel.Tracer("demo-api")

// Controller or Handler
func Hello(s HelloSevice) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		ctx, span := tracer.Start(c.UserContext(), "helloHandler", oteltrace.WithAttributes(attribute.String("layer", "handler")))
		defer span.End()

		Logger.Info("demo", zap.String("traceId",
			span.SpanContext().TraceID().String()))

		msg := s.doSth(ctx)
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}
}

// Service
type HelloSevice struct {
	r IRepository
}

func (s HelloSevice) doSth(ctx context.Context) string {
	ctx, span := tracer.Start(ctx, "HelloSevice.doSth")
	defer span.End()
	Logger.Info("HelloSevice", zap.String("traceId",
		span.SpanContext().TraceID().String()))
	return s.r.GetDataFromDb(ctx)
}

func NewService(r IRepository) HelloSevice {
	return HelloSevice{r: r}
}

// Repository
type IRepository interface {
	GetDataFromDb(ctx context.Context) string
}

type HelloRepository struct {
	Client *mongo.Client
}

type Demo struct {
	Message string `bson:"message,omitempty"`
}

func (r HelloRepository) GetDataFromDb(ctx context.Context) string {
	// TODO :: demo_collection
	c := r.Client.Database("test").Collection("demo_collection")
	var result Demo
	c.FindOne(context.TODO(), bson.D{}).Decode(&result)
	fmt.Println(result)
	_, span := tracer.Start(ctx, "HelloRepository.GetDataFromDb")
	defer span.End()

	Logger.Info("HelloSevice", zap.String("traceId",
		span.SpanContext().TraceID().String()))

	return result.Message
}
