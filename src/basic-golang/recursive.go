package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	
}

func factorial(angka int) int {
	if angka == 1 || angka == 0{
		return 1
	}
	return angka * factorial(angka-1)
}

