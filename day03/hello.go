package day03

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHelloRouter(app *fiber.App, s HelloSevice) {
	app.Get("/", Hello(s))
}

// Controller or Handler
func Hello(s HelloSevice) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		msg := s.doSth()
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}
}

// Service
type HelloSevice struct {
	r IRepository
}

func (s HelloSevice) doSth() string {
	return s.r.GetDataFromDb()
}

func NewService(r IRepository) HelloSevice {
	return HelloSevice{r: r}
}

// Repository
type IRepository interface {
	GetDataFromDb() string
}

type HelloRepository struct {
	Client *mongo.Client
}

type Demo struct{
	Message string  `bson:"message,omitempty"`
}

func (r HelloRepository) GetDataFromDb() string {
	// TODO :: demo_collection
	c := r.Client.Database("test").Collection("demo_collection")
	var result Demo
	c.FindOne(context.TODO(), bson.D{}).Decode(&result)
	fmt.Println(result)
	return result.Message
}
