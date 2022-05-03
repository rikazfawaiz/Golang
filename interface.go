package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHay(hasname HasName) {
	fmt.Println("Hi",hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (sato Animal) GetName() string {
	return sato.Name
}

func main() {
	name := Person{Name: "rikaz"}
	sayHay(name)
	name2 := Animal{Name: "ucing"}
	sayHay(name2)
}