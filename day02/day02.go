package day02

import "fmt"

type Hello struct{
	Name string
}

func (h Hello) String() string {
	return fmt.Sprintf("Name=%s", h.Name)
}

func NewHello(name string) Hello {	
	return Hello{ Name: name }
}