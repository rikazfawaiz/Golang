package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rikaz Fawaiz","Rikaz"))
	fmt.Println(strings.Split("Rikaz Fawaiz"," "))
	fmt.Println(strings.ToLower("Rikaz Fawaiz"))
	fmt.Println(strings.ToUpper("Rikaz Fawaiz"))
	fmt.Println(strings.Trim("   Rikaz Fawaiz   "," "))
	fmt.Println(strings.ReplaceAll("Rikaz Fawaiz","a","i"))
}