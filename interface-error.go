package main

import (
	"errors"
	"fmt"
)

func Pembagian (nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai/pembagi, nil
	}
}

func main() {
	nilai, err := Pembagian(100,5)
	if err == nil {
		fmt.Println("hasil",nilai)
	} else {
		fmt.Println("error", err.Error())
	}
}