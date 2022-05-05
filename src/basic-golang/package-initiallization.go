package main

import (
	_ "basic-golang/database"
	_ "fmt" 
	_ "basic-golang/other" // kalau package diimport tapi gak dipake kasih _ sebelum nama package (blank identifier)
)

func main() {
	// result := database.GetDatabase()
	// fmt.Println(result)
}