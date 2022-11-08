package main

import (
	"fmt"
	"io"
	"net/http"
)

// var wg sync.WaitGroup

func main() {
	_, err := fetch("https://jsonplaceholder.typicode.com/users")
	fmt.Println(err)
}

func fetch(url string) (string, error) {
	// Use package net/http
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), nil
}
