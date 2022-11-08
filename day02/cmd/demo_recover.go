package main

import (
	"fmt"
	"io/ioutil"
)

func doWithPanic() {
	err := recover()
	fmt.Println(err)
}

func main() {
	defer doWithPanic()
    // Read data from file
    b, err := ioutil.ReadFile("try_panic.go")
    if err != nil {
        panic("read file error")
    }

    fmt.Println(string(b))
}
