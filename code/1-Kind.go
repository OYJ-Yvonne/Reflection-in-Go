package main

import (
   "fmt"
   "reflect"
)

type User struct {
   Name string
   Age  int
}

type MyInt int

func main() {
   u := User{"xff", 19}

   var i MyInt = 42

   t := reflect.TypeOf(u)
   // reflect.Struct是一个reflect.Kind类型的常量，代表了结构体类型
   fmt.Println(t.Kind() == reflect.Struct) // 输出 true

   t = reflect.TypeOf(i)
   // 这里变量i的类型是我们自定义的类型MyInt，但是Type.Kind()方法返回的类型是reflect.Int
   fmt.Println(t.Kind() == reflect.Int) // 输出true

}

