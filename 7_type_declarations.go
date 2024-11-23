package main

import "fmt"

func main(){

	type id string

	var idDummy id = "111111"
	fmt.Println(idDummy)
	fmt.Println(id("111111"))
} 