package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    // Example: Calculate the factorial of 5
    number := 5
    result := factorial(number)
    fmt.Printf("Factorial of %d is %d\n", number, result)
}
