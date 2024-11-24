package main

import "fmt"

type validationError struct {
	Message string
}

func (v validationError) Error() string {
	return v.Message
}

type notFOundError struct {
	Message string
}

func (n notFOundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation Error"}
	}

	if id != "reza" {
		return &notFOundError{Message: "Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error ", validationErr.Message)
		} else if notFOundErr, ok := err.(*notFOundError); ok {
			fmt.Println("Not foun error ", notFOundErr.Message)
		} else {
			fmt.Println("Error Unknown", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
