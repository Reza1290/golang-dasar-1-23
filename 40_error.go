package main

import (
	"errors"
	"fmt"
) // there some package to handle error name errors or use interface
type error interface {
	Error() string
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	res, err := Pembagian(1, 0)

	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("error", err.Error())
	}

}
