package main
import "fmt"
func Ups() interface{}{
	return "Ups"
}
func main(){
	empty := Ups()

	fmt.Println(empty)
}