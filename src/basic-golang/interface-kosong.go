package main

import "fmt"

func Ups(i int) interface{} { //type apapun bisa
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	i := Ups(2)
	fmt.Println(i)
}