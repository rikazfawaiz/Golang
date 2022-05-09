package main

import (
	"fmt"
	"reflect"
)

//package reflect untuk melihat kode ketika program sedang running

type Sample struct {
	Name string `namatag:"nilaitag" max:10`
}

type Contoh struct {
	Name string `namatag:"nilaitag" max:10` 
	Desc string `namatag:"nilaitag" max:10`
}

//buat deteksi tag di function
func IsValid(data interface{}) bool{ 
	t := reflect.TypeOf(data)
	for i:=0; i < t.NumField(); i++{
		field := t.Field(i)
		if field.Tag.Get("namatag") == "nilaitag" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}

	}
	return true
}

func main() {
	sample := Sample{"Rikaz"}

	sampleType := reflect.TypeOf(sample)
	
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("namatag"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := Contoh{"",""}
	fmt.Println(IsValid(contoh))


}