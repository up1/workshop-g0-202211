package main

import (
	"fmt"
	"log"
	"net/http"
)

type conn struct{}

func main() {

	// Connect to db
	c := conn{}

	http.HandleFunc("/hello", helloHandler(c))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(c conn) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(c)
	}
}
