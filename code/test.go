//举例说明反射使用 TypeOf 和 ValueOf 来取得传入类型的属性字段于方法

package main

import (
    "fmt"
    "reflect"

    //"github.com/golang/protobuf/proto"
)

//定义一个用户结构体
type User struct {
    Id int
    Name string
    Age int
}

//为接口绑定方法
func (u User) Hello() {
    fmt.Println("Hello World.")
}

//定义一个可接受任何类型的函数（空接口的使用规则）
func Info(o interface{}) {
    t := reflect.TypeOf(o)    //获取接受到到接口到类型
    fmt.Println("Type:", t.Name())    //打印对应类型到名称(这是reflect中自带到)

    //Kind()方法是得到传入类型到返回类型；下面执行判断传入类型是否为一个结构体
    if k := t.Kind(); k != reflect.Struct {
        fmt.Println("传入的类型有误，请检查!")
        return
    }

    v := reflect.ValueOf(o)    //获取接受到到接口类型包含到内容(即其中到属性字段和方法)
    fmt.Println("Fields:")  //如何将其中到所有字段和内容打印出来呢？
    /**
    通过接口类型.NumField 获取当前类型所有字段个数
     */
    for i := 0; i < t.NumField(); i++ {
        f := t.Field(i)            //取得对应索引的字段
        val := v.Field(i).Interface()    //取得当前字段对应的内容
        fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
    }
    /**
    通过接口类型.NumMethod 获取当前类型所有方法的个数
     */
    fmt.Println("Method:")
    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)        //取得对应索引的方法
        fmt.Printf("%6s: %v\n", m.Name, m.Type)
    }
}


type FinalData struct {
    Rowkey_a         string  `json:"rowkey"`

}


func main()  {
    u := User{1, "OK", 12}
    Info(u)
    //Info(&u) 如果传入的是结构体的地址或指针(pointer-interface)，那么在Info函数中的Kind方法进行判断时就会被拦截返回


    //test
    //popDataRsp:= []FinalData{
    //    {
    //        Rowkey_a:"3434545",
    //    },
    //    {
    //        Rowkey_a:"43525346",
    //    },
    //    {
    //        Rowkey_a:"767867876",
    //    },
    //
    //}
    //finalData := make([]FinalData,0)
    //for _, popItem := range popDataRsp {
    //    var item FinalData
    //    item.Rowkey_a = popItem.Rowkey_a
    //    fmt.Println(item)
    //    finalData = append(finalData, item)
    //}
    //fmt.Println(finalData)
    //data, _ := json.Marshal(finalData)
    //
    //f, _ := os.OpenFile("./testData.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
    //
    //fmt.Println(string(data))
    //
    //if _, err := f.Write(data); err != nil {
    //    nlog.DfError("write fail:%+v", err)
    //}

    //test
    //finalData := make([]FinalData,0)
    //for _, popItem := range popDataRsp {
    //    var item FinalData
    //    item.Rowkey = popItem.GetRowkey()
    //    item.FeedsClick = popItem.GetMetric().GetFeedsClick()
    //    item.FeedsExpose = popItem.GetMetric().GetFeedsExpose()
    //    item.ShareCnt = popItem.GetMetric().GetShareCnt()
    //    item.PraiseCnt = popItem.GetMetric().GetPraiseCnt()
    //    item.BiuCnt = popItem.GetMetric().GetBiuCnt()
    //    item.JubaoCnt = popItem.GetMetric().GetJubaoCnt()
    //    item.CommentCnt = popItem.GetMetric().GetCommentCnt()
    //    item.CmtZanCnt = popItem.GetMetric().GetCmtZanCnt()
    //    item.FeedsValidNfbCnt = popItem.GetMetric().GetFeedsValidNfbCnt()
    //    item.ValidNfbRate = popItem.GetMetric().GetValidNfbRate()
    //    item.ContentQualityScore = popItem.GetMetric().GetContentQualityScore()
    //    finalData = append(finalData, item)
    //
    //}
    //data, err := json.Marshal(finalData)
    //f, _ := os.OpenFile("/usr/local/services/testData.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
    //if _, err := f.Write(data); err != nil {
    //    nlog.DfError("write fail:%+v", err)
    //}



    //test1 := []string{
    //   "11111",
    //   "22222",
    //   "33333",
    //   "44444",
    //   "55555",
    //   "66666",
    //   "77777",
    //   "88888",
    //   "99999",
    //}
    //test2 := make([]string,0)
    //test2 = test1
    //test3 := make([]string,0)
    //flag := 4
    //
    //if len(test1) < flag{
    //   flag = len(test1)
    //}
    //start := 0
    //end := flag
    //num := len(test1)/flag +1
    //if len(test1)%flag == 0 {
    //   num = len(test1)/flag
    //}else{
    //   num = len(test1)/flag +1
    //}
    //fmt.Println(num) // 输出true
    //
    //
    //for i:=0; i<num;i++{
    //   fmt.Println(test2[start:end]) // 输出true
    //   if i == 1{
    //       test3 = (test3)[0:0]  //从头来过，清空slice
    //   }
    //   test3 = append(test3,test2[start:end]...)
    //
    //   start += flag
    //   if i == len(test1)/flag-1 {
    //       end += len(test1)%flag
    //   }else{
    //       end += flag
    //   }
    //
    //}
    //fmt.Println(test3) // 输出true


    rowkeylist := []string{
      "11111",
      "22222",
      "33333",
      "44444",
      "55555",
      "66666",
      "77777",
      "88888",
      "99999",
    }
    rowkeylistFrag := make([]string,0)
    maxSize := 8

    //循环请求，每次100条rowkey
    for i := 0; i < len(rowkeylist); i+= maxSize {

        fmt.Println(i + maxSize) // 输出true
        fmt.Println(len(rowkeylist)) // 输出true
        if i + maxSize > len(rowkeylist) {
            fmt.Println(1) // 输出true
            rowkeylistFrag = rowkeylist[i : len(rowkeylist)]
        }else{
            fmt.Println(2) // 输出true

            rowkeylistFrag = rowkeylist[i : i+maxSize]
        }
        fmt.Println(rowkeylistFrag) // 输出true
        //err := getSingleExtraData(rowkeylistFrag, &totalExtraData, nc)

    }


}

// Type: User
// Fields:
//     Id: int = 1
//   Name: string = OK
//    Age: int = 12
// Method:
//  Hello: func(main.User)