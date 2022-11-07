package module_b

type ICustomer interface {
	SayHi() string
}

type Customer struct{
	Id int
	Name string
}

type CustomerV2 struct{
	Age int
	Customer
}

func NewCustomer() Customer {
	c2 := CustomerV2{}
	c2.Id = 1
	c2.Name = "xxx"
	c2.Age = 10

	return Customer{}
}

func (c Customer) SayHi() string {
	return "Hello from module B"
}