//如何通过反射得道结构当中匿名或者嵌入字段

package main

import (
    "fmt"
    "reflect"
)

//定义一个用户结构体
type User struct {
    Id int
    Name string
    Age int
}

type Manager struct {
    User    // 这是一个匿名字段，这个字段的名称也是User
    title string
}

func main() {
    m := Manager{User: User{1, "OK", 15}, title: "123"}
    t := reflect.TypeOf(m)

    //取得类型中的字段是否为匿名字段
    fmt.Printf("%+v\n", t.Field(0))

    //{User main.User  0 [ 0] true}
    //获取匿名类型中的字段，这里需要使用序号组，传入要取的切片即可

    fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))

    tchage := reflect.ValueOf(&m)    //想要修改和我们之前所说的传入值类型和指针类型是一致的，要想修改需要传入对应指针类型
    tchage.Elem().FieldByIndex([]int{0, 0}).SetInt(999) //传入指针需要通过 .Elem() 来取得对应的值内容，之后再想取哪个再继续使用序号组
    fmt.Println(tchage.Elem().FieldByName("title"))
    fmt.Println(tchage)
}

// {User  main.User  0 [0]  true}
// {Id  int  0 [0] false}
// 123
// &{{999 OK 15} 123}