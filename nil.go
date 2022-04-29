package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name":name,
		}
	}
}

func main() {
	var person = NewMap("rikaz")
	fmt.Println(person)
	//yg bisa pake nil -> interface, map, function, slice, pointer & channel sisanya default value, misal int jadi 0, string jadi ""
}