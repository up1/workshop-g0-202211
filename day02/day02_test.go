package day02_test

import (
	"day02"
	"fmt"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	h := day02.NewHello("somkiat")
	fmt.Println(h)
}