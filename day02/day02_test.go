package day02_test

import (
	"day02"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	h := day02.NewHello("somkiat")
	got := h.SayHi()
	if got != "Hello somkiat" {
		t.Errorf("Error mismatch with %s", got)
	}
}