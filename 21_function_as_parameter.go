package main

import "fmt"
// type declaration func
type Filter func(string) string // use this to simplify

func sayHelloWithFilter(name string, filter func(string) string) { 
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Reza", spamFilter)

	filter:= spamFilter

	sayHelloWithFilter("Anjing", filter)
}

