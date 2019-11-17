//通过反射对方法等动态调用

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
 
 func (u User) FuncHasArgs(name string, age int) {
	fmt.Println("FuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
 }
 
 func (u User) FuncNoArgs() {
	fmt.Println("FuncNoArgs")
 }
 
 func main(){
 
	user := User{1, "test", 13}

	//正常调用
	user.FuncHasArgs("jack", 20) 

	//反射调用
	uValue := reflect.ValueOf(user)

	uType  := reflect.TypeOf(user)
	
	m1 := uValue.MethodByName("FuncHasArgs") //IsValid
	
	m2 := uValue.MethodByName("FuncNoArgs")  
	
	m , ok := uType.MethodByName("FuncNoArgs") //return two values


	args1 := []reflect.Value{reflect.ValueOf("oyj"), reflect.ValueOf(3)}
	m1.Call(args1)

	args2 := make([]reflect.Value,0)
	m2.Call(args2)

	
	//output
	fmt.Printf("m1:%#v\n",m1)
	
	fmt.Printf("m2:%#v\n",m2)
	
	fmt.Printf("m:%#v, isfound:%v\n",m, ok)
	
 }
 