package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	reza := Man{"Reza"}
	reza.Married() // not yet hehe, i wrote this code infront of my future wife :D

	fmt.Print(reza)
}
