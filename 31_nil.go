package main

import "fmt"

// nil only for map, slice, interface, function, channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Reza")

	if data == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(data["name"])
	}
}
