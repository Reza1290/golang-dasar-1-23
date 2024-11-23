package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Test"
	names[1] = "Test2"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{1, 2, 3}

	fmt.Println(values[0])
	fmt.Println(values[1])

	// infinity buffer?

	var inf = [...]int{1, 23, 4, 5, 5, 5, 1, 1}

	fmt.Println(len(inf))
}
