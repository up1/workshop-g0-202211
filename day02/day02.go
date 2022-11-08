package day02

import "fmt"

// Hello struct
type Hello struct {
	Name string
}

// To String
func (h Hello) String() string {
	return fmt.Sprintf("Name=%s", h.Name)
}

// Create new Hello
func NewHello(name string) Hello {
	return Hello{Name: name}
}

// SayHi with name
func (h Hello) SayHi() string {
	return "Hello " + h.Name
}
