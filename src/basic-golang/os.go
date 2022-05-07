package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument",args)

	//go run os.go rikaz fawaiz
	// fmt.Println(args[1]) //rikaz
	// fmt.Println(args[2]) //fawaiz

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname",hostname)
	} else {
		fmt.Println("Error",err.Error())
	}
	

}