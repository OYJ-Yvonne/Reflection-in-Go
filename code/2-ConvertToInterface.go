package main

import (
    "fmt"
    "reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
 }
 
 func (u User) ReflectCallFunc() {
	fmt.Println("reflect learn")
 }

 
func main() {
 user := User{1, "test", 13}
 
 var i interface{}

 i = user
 
 uValue := reflect.ValueOf(i)
 
 uType  := reflect.TypeOf(i)

 fmt.Printf("Interface var:%v\n",uValue.Interface()) //转换为interface类型,unpack uValue.Interface().(User)

 for i := 0; i < uType.NumField(); i++ { //获取field信息
 
	field := uType.Field(i)

	value := uValue.Field(i).Interface()

	fmt.Printf("%s: %v = %+v\n", field.Name, field.Type, value)
 
 }


}
