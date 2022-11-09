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
	r HelloRepository
}

func (s HelloSevice) doSth() string {
	return s.r.getDataFromDb()
}

func NewService(r HelloRepository) HelloSevice {
	return HelloSevice{r: r}
}

// Repository
type HelloRepository struct {
}

func (r HelloRepository) getDataFromDb() string {
	panic("Under construction !!")
}
