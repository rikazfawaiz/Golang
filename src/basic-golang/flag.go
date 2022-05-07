package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host","hostname","Put your database host")
	user := flag.String("user","root","Put your database username")
	pass := flag.String("pass","root","Put your database password")

	flag.Parse()
	fmt.Println(*host, *user, *pass) //go run flag.go -host=hostname -user=root -pass=root --> hostname, root, root
}