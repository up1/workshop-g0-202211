package day02

type Hello struct{
	Name string
}

func NewHello(name string) Hello {	
	return Hello{ Name: name }
}