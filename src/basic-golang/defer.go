//defer menjalankan fungsi walaupun program error
package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("run applicarion")
}

func main() {
	runApplication()
}
