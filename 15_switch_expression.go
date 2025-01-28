package main

import "fmt"

func main() {

	name := "test"

	switch name {
	case "test":
		fmt.Println("benar")
	case "test2":
		fmt.Println("kurang")
	default:
		fmt.Println("tetot")
	}

	switch len := len(name); len > 1 {
	case true:
		fmt.Println("To long")
	case false:
		fmt.Println("To short")
	}

	switch {
	case name == "test":
		fmt.Println("benar")
	case name == "test2":
		fmt.Println("idk")
	}
}
