package main

import (
    "fmt"
    "reflect"
)

type T struct {

	A int
 
	B string
 
 }
 
 func main() {
	t := T{23, "skidoo"}
	
	s := reflect.ValueOf(&t).Elem()

	fmt.Println("canSet:",s.CanSet())

	typeOfT := s.Type()
	
	for i := 0; i < s.NumField(); i++ {
		//f := s.Field(i)
		if typeOfT.Field(i).Name == "A" {
			s.Field(i).SetInt(20)
		}
		//fmt.Printf("f = %v\n", f.Interface())

	}
	fmt.Printf("after change1: %v\n", s)


	//FieldByName()
	f := s.FieldByName("B")
	if !f.IsValid() {
		fmt.Println("没有B对应属性字段")
	}
	f.SetString("test3333")
	fmt.Println("after change3: ",s.Interface())


 }