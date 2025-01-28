package main

import "fmt"

type Customer struct {
	Name, Alamat string
	Usia         int
}

func main() {
	var me Customer
	me.Name = "test"
	me.Alamat = "jalan"
	me.Usia = 18

	fmt.Println(me)
	fmt.Println(me.Name)

	budi := Customer{
		Name:   "Budi",
		Alamat: "Jalan",
		Usia:   24,
	}

	fmt.Println(budi)

	doe := Customer{"Doe", "makan", 10}

	fmt.Println(doe)
}
