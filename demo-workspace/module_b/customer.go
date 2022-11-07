package module_b

type Customer struct{
	Id int
	Name string
}

func NewCustomer() Customer {
	return Customer{}
}

func (c Customer) SayHi() string {
	return "Hello from module B"
}