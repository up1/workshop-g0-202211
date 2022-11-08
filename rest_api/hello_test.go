package demo_test

import (
	"demo"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessWithHello(t *testing.T) {
	// Setup router
	app := demo.NewRouter()
	demo.NewHello(app, nil)

	// Call Target endpoint
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, _ := app.Test(req, -1)

	// Assert status code
	assert.Equal(t, 200, res.StatusCode)
	// Assert body
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, nil, err)
	assert.Contains(t, string(body), "Hello")
}
