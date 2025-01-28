package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "test",
		"address": "sby",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "test"
	book["ups"] = "Salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
