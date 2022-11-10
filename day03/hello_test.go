package day03_test

import (
	"context"
	"day03"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccesswithHelloHandler(t *testing.T) {
	app := day03.NewFiberRouter()
	// init dependency
	client := day03.NewMongoClient("mongodb://user:pass@128.199.205.113:27017")
	r := day03.HelloRepository{Client: client}
	s := day03.NewService(r)
	day03.NewHelloRouter(app, s)
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

type StubRepository struct{}

func (r StubRepository) GetDataFromDb(context.Context) string {
	return "From stub"
}

func TestSuccesswithHelloHandlerWithStub(t *testing.T) {
	app := day03.NewFiberRouter()
	// init dependency
	r := StubRepository{}
	s := day03.NewService(r)
	day03.NewHelloRouter(app, s)
	// Call Target endpoint
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, _ := app.Test(req, -1)

	// Assert status code
	assert.Equal(t, 200, res.StatusCode)
	// Assert body
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, nil, err)
	assert.Contains(t, string(body), "From stub")
}
