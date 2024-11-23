package main

import "fmt"

type Sigma func(string) bool

func sigmaUser(name string, sigma Sigma) {
	if sigma(name){
		fmt.Println("You are sigma ", name)
	}else{
		fmt.Println("Not Sigma ", name)
	}
}


func main() {
	sigmas:= func(name string) bool {
		return name == "Reza"
	}

	sigmaUser("Reza", sigmas)
}
