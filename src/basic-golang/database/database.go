package database

import "fmt"

var connection string

func init() { //dirunning paling pertama ketika package diakses
	connection = "MySQL"
	fmt.Println("Test Init")
}

func GetDatabase() string {
	return connection
}