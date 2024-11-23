package main

import "fmt"



func dummyFunc(error bool){
	defer helloWorld() // call when the func end

	if error {
		panic("ERROR")
	}
	fmt.Println("me Dummy")

	
}

func helloWorld(){
	fmt.Println("Hello World")
	message := recover()
	fmt.Println(message)
}

func main() {
	dummyFunc(true)
}
