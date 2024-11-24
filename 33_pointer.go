package main
import "fmt"

type Adress struct{
	Kota, Provinsi, Negara string
}

func main(){
	var address1 Adress = Adress{"Surabaya","Jawa Timur","Indonesia"}
	var address2 *Adress = &address1

	address2.Kota = "Bekasi"

	fmt.Println(address1)
	fmt.Println(address2)

	
}