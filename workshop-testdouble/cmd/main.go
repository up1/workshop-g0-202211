package main

import (
	"fmt"
	"mock"
)

func main() {

	websites := []string{
		"http://www.google.com",
		"http://www.somkiat.cc",
	}

	results := mock.CheckAllWebs(websites)

	for url, ok := range results {
		fmt.Println(url, " : ", ok)
	}

}
