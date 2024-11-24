package main
import "fmt"

type User struct{
	Name, Alamat string
	id int
}

func (user User) sayHello(){
	fmt.Println("Hello, My Name is ", user.Name)
}

func main(){
	reza:= User{"Reza","Jalan",20}
	reza.sayHello()
}