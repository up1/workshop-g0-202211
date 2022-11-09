package main

import (
	"fmt"
	"log"
	"net/http"
)

type conn struct{}

type hello struct {
	c conn
}

func (h hello) helloHandler3(w http.ResponseWriter, req *http.Request) {
	fmt.Println(h.c)
}

func main() {

	// Connect to db
	c := conn{}
	h := hello{c: c}
	http.HandleFunc("/hello", h.helloHandler3)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler2(c conn) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(c)
	}
}
