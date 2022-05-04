//panic untuk menghentikan program yang sedang berjalan
package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")	
}

func main() {
	runApp(false)
	fmt.Println("test app masih jalan")
}