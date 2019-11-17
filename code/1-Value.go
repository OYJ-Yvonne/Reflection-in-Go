
package main

import (
    "fmt"
    "reflect"
)

type User1 struct {
    Name string
    Age  int
}

func (u User1) test1() {
    fmt.Println("test")
}

type User2 struct {
    Name string
    Age  int
}

func (u User2) test1() {
    fmt.Println("test")
}

type MyInt1 int
type MyInt2 int
func main() {
    //u1 := User1{"xff", 19}
    //u2 := User2{"xff", 19}
    //
    //t1 := reflect.TypeOf(u1)
    //t2 := reflect.TypeOf(u2)
    //
    //fmt.Println(t1)
    //fmt.Println(t2)
    //
    //
    //v1 := reflect.ValueOf(u1)
    //v2 := reflect.ValueOf(u2)
    //
    //fmt.Printf("%+v\n",v1)
    //fmt.Printf("%+v\n",v2)
    //
    //fmt.Printf("%+v\n",v1.Interface())
    //
    //fmt.Printf("%+v\n",v2.Interface())
    //
    //fmt.Println(v1==v2)

    var x1 MyInt1 = 1
    var x2 MyInt2 = 1

    v1 := reflect.ValueOf(x1)
    v2 := reflect.ValueOf(x2)

    fmt.Println(v1==v2)  //false

}
