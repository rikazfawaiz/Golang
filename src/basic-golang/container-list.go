package main

import (
	"container/list"
	"fmt"
)

func main() {
	//double link list
	data := list.New()
	data.PushBack("Rikaz") //menambahkan paling belakang
	data.PushFront("Asep") //menambahkan paling depan
	data.PushBack("Fawaiz")
	data.PushBack("Haerul")
	data.PushBack("Afgani")

	fmt.Println("Front :",data.Front().Value)
	fmt.Println("Back :",data.Back().Value)
	
	fmt.Println("Next :",data.Front().Next().Next().Next().Value) //ngambil yg next"" nya
	fmt.Println("Prev :",data.Front().Prev()) //ngambil data sebelum front

	//depan ke belakang
	// for e := data.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	//belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}