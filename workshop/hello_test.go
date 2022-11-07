package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestXXX(t *testing.T) {
	want := "Hello somkiat"
	got := sayHi("somkiat")
	if got != want {
		t.Errorf("Got %s, Want %s", got, want)
	}
	assert.Equal(t, want, got)
}
