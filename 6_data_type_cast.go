package main

import "fmt"

func main(){

	var value32 int32 = 32582
	var value64 int64 = int64(value32)

	var value16 int16 = int16(value64)

	fmt.Println(value16)
	fmt.Println(value32)
	fmt.Println(value64)

	name := "MyName"

	var e uint8 = name[0] //ASCII UINT8 = learn more googling
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}	