package main
import "fmt"

type Alamat struct{
	Kota, Provinsi, Negara string
}

func main(){
	var address1 Alamat = Alamat{"Surabaya","Jawa Timur","Indonesia"}
	var address2 *Alamat = &address1

	address2.Kota = "Bekasi"

	fmt.Println(address1)
	fmt.Println(address2)
	
	address2 = &Alamat{"Bandung","Jawa Barat","Indonesia"}
	
	fmt.Println(address1)
	fmt.Println(address2)
	
	*address2 = Alamat{"Jakarta","DKI","Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}