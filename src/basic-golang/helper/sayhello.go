package helper

import "fmt"

//Akses Modifier, Huruf Besar akses global (diluar package), huruf kecil akses private diawal nama func & variable
var version = "1.1.1"
var Application = "Golang"

func SayHello(name string) {
	fmt.Println("Hello",name)
}

func sayGoodBye(name string) {
	fmt.Println("Goodbye",name)
}