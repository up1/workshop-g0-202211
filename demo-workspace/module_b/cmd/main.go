package main

import (
	"fmt"
	"module_b"
)

func main(){
	c := module_b.NewCustomer()
	fmt.Println(c.SayHi())
}