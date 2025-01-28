package main

import "fmt"

func main() {

	test()
	funcWithParam(1)
	funcWithParam(2)
	fmt.Println(funcWithReturn(1,4))

	val, str := funcWithReturnMultipleval(1,10) //same destruct like js/py

	fmt.Println(val)
	fmt.Println(str)
}

func test() {
	fmt.Println("test")
}

func funcWithParam(val int){
	fmt.Println(val)
}

func funcWithReturn(a int, b int) int {
	return a + b
}

func funcWithReturnMultipleval(a int, b int) (int, string){ // actually we can add name return here like (val int) or (a, b, c string)
	return a + b , "bruh"
}