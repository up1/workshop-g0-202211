package day03_test

import (
	"day03"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSuccesswithHelloHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/", day03.Hello)
	// Call Target endpoint
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, _ := app.Test(req, -1)

	// Assert status code
	assert.Equal(t, 200, res.StatusCode)
}
