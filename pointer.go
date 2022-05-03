package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Garut", "Jawa Barat", "Indonesia"}
	address2 := &address1 //passing by reference, nilai address1 berubah jika address2 berubah
	address3 := &address1
	address4 := new(Address) //buat variabel pointer kosong

	address2.City = "Bandung"
	// address2 = &Address{"Tasik", "Jawa Barat", "Indonesia"} //hanya mengubah address2 tanpa mengubah variabel yg di referencenya / membuat variabel pointer baru
	*address3 = Address{"Bekasi", "Jawa Barat", "Indonesia"} //mengubah nilai variabel yg di referencenya ke dirinya

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}