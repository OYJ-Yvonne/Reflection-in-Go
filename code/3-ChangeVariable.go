package main

import (
    "fmt"
    "reflect"
)

func main() {
	var x float64 = 3.4

	p := reflect.ValueOf(&x) // Note: take the address of x.

	fmt.Println("type of p:", p.Type())

	fmt.Println("settability of p:", p.CanSet())

	// print:

	// type of p: *float64

	// settability of p: false 

	// p不可set,p指向的内容可set,p指向的内容即*p

	// reflect.Value 的Elem方法，可以获取value 指向的内容

	v := p.Elem()

	fmt.Println("settability of v:", v.CanSet()) //settability of v: true

	v.SetFloat(7.1)

	fmt.Println(v.Interface()) //7.1

	fmt.Println(x)             //7.1
}
