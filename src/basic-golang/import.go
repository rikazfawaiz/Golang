package main

import (
	"basic-golang/helper"
	"fmt"
)

func main() {
	for counter := 1; counter < 10; counter++ {
		helper.SayHello("rikaz")
	}
	fmt.Println(helper.Application)
}