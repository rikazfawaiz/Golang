//panic untuk menghentikan program yang sedang berjalan
package main

import "fmt"

func endApp() {
	message := recover() // menangkap pesan error dari panic dan tetap menjalankan program walau error
	if message != nil {
		fmt.Println("Pesan Error :", message)
	}
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
	runApp(true)
	fmt.Println("test app masih jalan")
}
