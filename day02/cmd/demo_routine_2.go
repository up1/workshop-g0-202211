package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go fetch("https://www.google.com")

	wg.Add(1)
	go fetch("https://jsonplaceholder.typicode.com/users", )

	wg.Wait()
}

func fetch(url string) (string, error) {
	defer wg.Done()
	// Use package net/http
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println("Done with ", url)
	return string(body), nil
}
