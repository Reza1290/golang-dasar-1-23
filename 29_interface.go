package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHelo(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"reza"}
	sayHelo(person)
	animal := Animal{"Budi"}
	sayHelo(animal)
}
