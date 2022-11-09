package day03

import "github.com/gofiber/fiber/v2"

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
}

func (r HelloRepository) GetDataFromDb() string {
	panic("Under construction !!")
}
