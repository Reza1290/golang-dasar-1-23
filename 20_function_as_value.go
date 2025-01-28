package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {

	bye := getGoodBye
	diffvar := getGoodBye

	fmt.Println(bye("Reza"))
	fmt.Println(diffvar("Reza"))

}
