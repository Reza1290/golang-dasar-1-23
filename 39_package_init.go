package main

import (
	"fmt"
	"hello-world/database"
)

// some package will gone if we not use it, so use blank identifier instead -> _ "hello-world/test"

func main() {
	fmt.Println(database.GetDatabase())
}
