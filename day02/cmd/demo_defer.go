package main

import (
	"fmt"
)


func main() {
	for i:=0 ;i<5; i++ {
		defer doSth(i)
	}
	fmt.Println("Finish")
}

func doSth(i int) {
	fmt.Println("Call doSth() ", i)
}