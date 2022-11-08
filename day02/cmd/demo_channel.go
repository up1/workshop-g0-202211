package main

import "fmt"

func main(){
	fmt.Println("Start")
	c := make(chan string)
	go slow2(c)
	go slow2(c)
	fmt.Println("Finish")
	for i:=0; i<2; i++{
		<-c
	}
}

func slow2(out chan<- string) {
	fmt.Println("Called slow()")
	out <- "OK"
}