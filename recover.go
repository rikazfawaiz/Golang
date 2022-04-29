package main

import "fmt"

func endApp() {
	message := recover() // menangkap pesan error dari panic dan tetap menjalankan program walau error// recover simpen di differ func
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


/*
nnnmm
m,mmm
*/

//vvvvv