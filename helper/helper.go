package helper

// the name must same name as parent dir until root dir
// case : a/b/c
var version = "1.0.0" 
var Application = "golang"

// Method public 
func SayHello(name string) string {
	return "Hello" + name
}

// cant be called outside package start with smallcase, private
func sayHey() string {
	return "Hi"
}
