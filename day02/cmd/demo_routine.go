package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start main")
	wg.Add(1)
	go slow()
	wg.Add(1)
	go slow()
	wg.Wait()
	fmt.Println("Finish main")
}

func slow() {
	defer wg.Done() 
	fmt.Println("Called slow()")
}