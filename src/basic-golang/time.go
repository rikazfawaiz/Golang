package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Date())

	utc := time.Date(2022, 5, 9, 16, 59, 0, 0, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" //format default di golang, bukan yymmdd dsb
	parse, _ := time.Parse(layout,"2022-05-09")
	fmt.Println(parse)
}