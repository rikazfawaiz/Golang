package main

import (
	"fmt"
)

func random() interface{} {
	return 1.7
}

func main() {
	result := random()
	// fmt.Println(result.(string))
	// fmt.Println(result.(int))

	switch val := result.(type) {
	case string:
		fmt.Println("value string", val)
	case int:
		fmt.Println("value int", val)
	default:
		fmt.Println("Unknown")
	}
}