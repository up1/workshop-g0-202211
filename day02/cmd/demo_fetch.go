package main

import (
	"fmt"
	"net/http"
)

type XXX func(i int) string

type result struct {
	string
	bool
}

func main(){
	urls := []string{
		"https://www.google.com",
		"https://www.somkiat.ccc",
	}

	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, FetchData(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		fmt.Println(r.string, ":", r.bool)
	}
}

func FetchData(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	return response.StatusCode == http.StatusOK
}