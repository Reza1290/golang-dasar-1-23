package main

import "fmt"

func main(){
	const firstName string = "Reza"
	const lastName = "1290"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		test = "test"
		test2 = "test"
	)

	fmt.Println(test)
	fmt.Println(test2)
}