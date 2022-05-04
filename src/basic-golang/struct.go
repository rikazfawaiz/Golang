package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct function
func (customer Customer) sayHay() {
	fmt.Println("hi",customer.Name)
}

func main() {
	var rikaz Customer
	rikaz.Name = "Rikaz"
	rikaz.Address = "Garut"
	rikaz.Age = 20

	rikaz.sayHay()

	// rikaz2 := Customer {
	// 	Name: "Rikaz",
	// 	Address: "Garut",
	// 	Age: 20,
	// }

	// rikaz3 := Customer {"Rikaz", "Garut", 30}

	// fmt.Println(rikaz)
	// fmt.Println(rikaz2)
	// fmt.Println(rikaz3)
}