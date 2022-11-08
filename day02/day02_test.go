package day02_test

import (
	"day02"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWithName(t *testing.T) {
	h := day02.NewHello("somkiat")
	got := h.SayHi()
	assert.Equal(t, "Hello somkiat", got)
}