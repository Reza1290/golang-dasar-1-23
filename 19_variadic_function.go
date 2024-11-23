package main

import "fmt"

func sumAll(val ...int) int {
	total := 0

	for _, number := range val{
		total+=number
	}
	return total
}

func main() {
	total:= sumAll(10,2,4,5,9,6,7)
	
	fmt.Println(total)

	// slice
	num := []int{10,10,10,10,10}
	total = sumAll(num...)

	fmt.Println(total)
}
