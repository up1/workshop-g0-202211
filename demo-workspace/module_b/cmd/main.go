package main

import (
	"fmt"
	"module_b"
)

func main(){
	c := module_b.NewCustomer()
	c.Id = 12345
	fmt.Println(c.SayHi())
	fmt.Println(c.Id)
}