package main


func main() {
	
	numbers := make(map[string]int)

    numbers["one"] = 1
	_, found := numbers["two"]
    if !found {
		// Found data by key=two
	}


}