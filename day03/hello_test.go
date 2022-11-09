package day03_test

import (
	"day03"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccesswithHelloHandler(t *testing.T) {
	app := day03.NewFiberRouter()
	day03.NewHelloRouter(app)
	// Call Target endpoint
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, _ := app.Test(req, -1)

	// Assert status code
	assert.Equal(t, 200, res.StatusCode)
}
