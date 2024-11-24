package main
import "fmt"

type Tempat struct{
	Kota, Provinsi, Negara string
}

func main(){
	var tempat1 *Tempat= new(Tempat)
	var tempat2 *Tempat = tempat1

	tempat2.Negara = "Indonesia"

	fmt.Println(tempat1)
	fmt.Println(tempat2)
}