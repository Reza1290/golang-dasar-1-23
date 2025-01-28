package main

import "fmt"

type Place struct {
	Kota, Kabupaten, Negara string
}

func ChangePlaceToIndonesia(place *Place) {
	place.Negara = "Indonesia"
}

func main() {
	var place *Place = &Place{"Subang", "Jawa Barat", ""}
	ChangePlaceToIndonesia(place)

	fmt.Println(place)

}
