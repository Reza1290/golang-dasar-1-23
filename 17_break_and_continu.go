package main

import "fmt"

func main() {

	counter := 0

	for counter <= 10 {
		if counter == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(counter)
		counter++
	}
	for counter <= 10 {
		counter++
		if counter % 2 == 0 {
			fmt.Println("break")
			continue
		}
		fmt.Println(counter)
	}
}
