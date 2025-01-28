package main

import "fmt"

func main() {

	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	slice := array[0:2]

	fmt.Println("Initial Slice:")
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	// Appending to a slice
	slice = append(slice, 9, 10)
	fmt.Println("After Appending:")
	fmt.Println(slice)

	// Making a new slice
	newSlice := make([]int, 5)
	fmt.Println("New Slice with make:")
	fmt.Println(newSlice)

	// Copying a slice
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("Copied Slice:")
	fmt.Println(copiedSlice)

	// slice[low:high]
	// slice[low:] start - last index
	// slice[:high] 0 - (end - 1)
	// slice[:]
}
