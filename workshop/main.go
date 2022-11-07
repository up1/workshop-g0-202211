package main

import "fmt"

func main() {
	r1 := sayHi("somkiat")
	fmt.Println(r1)

	if _, err := withError(1); err != nil {
		fmt.Println(err)
	}
}

func withError(code int) (string, error) {
	if code == 0 {
		return "", nil
	}
	return "", fmt.Errorf("Have error")
}

func sayHi(name string) string {
	return "Hello " + name
}

func sayHi2(name string) (result string) {
	result = fmt.Sprintf("Hello %s", name)
	return
}
