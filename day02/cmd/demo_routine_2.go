package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var urls [2]string
	urls[0] = "https://www.google.com"
	urls[1] = "https://jsonplaceholder.typicode.com/users"

	for i:=0 ; i<len(urls); i++ {
		wg.Add(1)
		go fetch(urls[i])
	}

	for _, url := range urls {
		wg.Add(1)
		go fetch(url)
	}

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
