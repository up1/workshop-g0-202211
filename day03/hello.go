package day03

import "github.com/gofiber/fiber/v2"

func NewHelloRouter(app *fiber.App) {
	app.Get("/", Hello)
}
// Controller or Handler
func Hello(c *fiber.Ctx) error {
	r := helloRepository{}
	s := NewService(r)
	msg := s.doSth()

	return c.JSON(fiber.Map{
		"message": msg,
	})
}

// Service
type helloSevice struct {
	r helloRepository
}

func (s helloSevice) doSth() string {
	return s.r.getDataFromDb()
}

func NewService(r helloRepository) helloSevice {
	return helloSevice{r: r}
}


// Repository
type helloRepository struct {
}

func (r helloRepository) getDataFromDb() string {
	panic("Under construction !!")
}



