package module_a

import "module_b"

func SayHi() string {
	module_b.NewCustomer().SayHi()
	return "Hello from module A"
}

