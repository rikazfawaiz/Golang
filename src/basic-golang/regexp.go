package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z]|[A-Z])o`)

	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("elo"))
	fmt.Println(regex.MatchString("edp"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("edo edi eko ewo edo eki poi qwe", -1)
	fmt.Println(search)
}