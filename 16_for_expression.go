package main
import "fmt"
func main(){

	counter := 0

	for counter <= 10{
		fmt.Println(counter)
		counter++
	}

	for i:=0; i < 10; i++ {
		fmt.Println(i)
	}

	names := []string{"Reza","1290"}
	for i, x := range names{
		fmt.Println("index", i, "=", x)
	}
}
