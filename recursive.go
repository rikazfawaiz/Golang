package main

import "fmt"

func main() {
	// factorial(5)
	fibonacci()
}

func factorial(angka int) int {
	fmt.Println(angka)
	if angka == 1 || angka == 0{
		return 1
	}
	return angka * factorial(angka-1)
}

func fibonacci() {


}