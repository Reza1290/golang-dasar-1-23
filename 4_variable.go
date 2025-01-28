package main

import "fmt"

func main(){
	var name string

	name = "Reza1290"

	fmt.Println(name)

	name = "Hehe"

	fmt.Println(name)

	var last_name = "1290"

	fmt.Println(last_name)

	var first_name string = "Reza"
	fmt.Println(first_name)

	// var is not neccesarry but we use different way (:=)
	full_name := "Reza1290"
	fmt.Println(full_name)


	// multiple variable

	var (
		firstName = "Reza"
		lastName = "1290"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	
}