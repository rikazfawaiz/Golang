package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	//link list ring, nyambung muter
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "data " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})

}