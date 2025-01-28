package main

import "fmt"

func random() interface{} { // any
	return "OK"
}

func main() {
	var result any = random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // will panic cuz its not int
	// fmt.Println(resultInt)

	// try avoid panic
	switch value := result.(type) {
		case string:
			fmt.Println(value, " is String")
		case int:
			fmt.Println(value, " is int")
		default:
			fmt.Println("Unk")
	}
}
