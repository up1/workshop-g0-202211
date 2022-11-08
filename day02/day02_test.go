package day02_test

import (
	"day02"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWithName(t *testing.T) {
	h := day02.NewHello("somkiat")
	got := h.SayHi()
	assert.Equal(t, "Hello somkiat", got)
}

func ExampleHello_SayHi() {
	h := day02.NewHello("somkiat")
	got := h.SayHi()
	fmt.Println(got)
	// Output: Hello somkiat
}

func BenchmarkV1SayHi(b *testing.B) {
	h := day02.NewHello("somkiat")
    for i := 0; i < b.N; i++ {
		h.SayHi()
    }
}

func BenchmarkV2SayHi(b *testing.B) {
	h := day02.NewHello("somkiat")
    for i := 0; i < b.N; i++ {
		h.SayHi()
    }
}