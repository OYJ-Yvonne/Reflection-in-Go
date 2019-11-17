# Reflection in Go

jingouyang

---
### `Table Of Contents`
1. Definition Of Reflection

2. Types And Interfaces

3. <font color="#DAA520">Three Laws Of Reflection</font>

4. Applications Of Reflection

5. Characteristics Of Reflection

6. Improve Reflection Performance

7. Reference Link


---
### ☝ ️`Definition Of Reflection`

* <p align="left">Reflection is the ability to dynamically acquire program structure information (meta-information) at runtime.</p>

<p align="left">反射(reflection)是指在运行时，动态获取程序结构信息（元信息）的一种能力.</p>

<p align="left">借助某种手段检查自己结构的一种能力，通常就是借助编程语言中定义的各种类型（types）.</p>


---
### ☝ ️`Types And Interfaces`
<font color="#DAA520">Reflection is built on the type system</font>

```go
package main

type MyInt int

func main() {
    var i1 int = 1
    var i2 MyInt = 2

    i1 = i2 //cannot use i2 (type MyInt) as type int in assignment
}

```
<p align="left">They cannot be assigned to each other without casting type.</p>


---
### ☝ ️`Types And Interfaces`
* <font color="#DAA520">An interface</font> represents a set of determined method sets
* <font color="#DAA520">An interface variable</font> can store any specific value, as long as the type to which the specific value belongs implements all methods of the interface.

* 一个接口表示用一组确定的方法的集合。
* 一个接口变量能存储任意的具体值，只要这个具体值所属的类型实现了这个接口的所有方法。


---
### ☝ ️`Types And Interfaces`
<font color="#DAA520">interface(io.Reader and io.Writer)</font>
```go
// Reader is the interface that wraps the basic Read method.
type Reader interface {
    Read(p []byte) (n int, err error)
}
// Writer is the interface that wraps the basic Write method.
type Writer interface {
    Write(p []byte) (n int, err error)
}
var r io.Reader
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)
// and so on
```

---
### ☝ ️`Types And Interfaces`
<font color="#DAA520">Empty Interface</font>
* interface{} : method collection is empty and can hold any value
* pair<value, type>
* Principle : Using it to implement the reflection mechanism


---
### ☝ `Three Laws Of Reflection`
<p align="left"><font color="#DAA520">1. Reflection goes from interface value to reflection object.</font></p>

* 反射可以获取interface类型变量的具体信息

```go
// pair<value, concrete type>
// a new Value initialized to the concrete value

// ValueOf(nil) returns the zero Value.
func ValueOf(i interface{}) Value {}
// reflection Type that represents the dynamic type of i.

// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type {}

```

---
<p align="left"><font color="#DAA520">1. Reflection goes from interface value to reflection object</font></p>
<p align="center">TypeOf</p>

```go
func main() {

    var x float64 = 3.4
    
    fmt.Println("type:", reflect.TypeOf(x))
}
//type: float64
```
<font size="5">x首先被保存到一个空接口中，这个空接口然后被作为参数传递。</br>
reflect.Typeof会把这个空接口拆包（unpack）恢复出类型信息。</font>


---
<p align="left"><font color="#DAA520">1. Reflection goes from interface value to reflection object</font></p>
<p align="center">ValueOf</p>

```go
var x float64 = 3.4

fmt.Println("value:", reflect.ValueOf(x))
//Valueof方法会返回一个Value类型的对象

//value: <float64 Value>
```
```
Value.Type() // concrete type
Value.Kind() // underlying type
Type.Kind() // underlying type
```

[code-->Kind.go]()


---
### ☝ `Three Laws Of Reflection`

<p align="left"><font color="#DAA520">2. Reflection goes from reflection object to interface value.</font></p>

* 反射可以将反射对象转换成接口变量

```go
//<reflect.Value,reflect.Type>
reflect.ValueOf(i).Interface()
```
```go
//Interface方法把类型和值的信息打包成一个接口表示并且返回结果
// Interface returns v's value as an interface{}.
func (v Value) Interface() interface{}
```
[code-->ConvertToInterface]()


---
### ☝ `Three Laws Of Reflection`

<p align="left"><font color="#DAA520">3. To modify a object, the value must be settable.</font></p>

* 通过反射修改变量

```go
var x float64 = 3.4

v := reflect.ValueOf(x)

//print: settability of v: false
fmt.Println("settability of v:", v.CanSet()) 

// Error: will panic.
v.SetFloat(7.1) 

```
<p align="left">The reason：ValueOf is a copy of x, not x itself</p>

[code-->:ChangeVariable.go]()


---
<p align="left"><font color="#DAA520">3. To modify a object, the value must be settable.</font></p>

* 通过反射操作struct实例

```
//Traversing the structure variable to get the domain

Type.NumField()

Type.Field() {name, type}

Type.FieldByName('')
```
[code-->ChangeStruct.go]()



---
* 通过反射动态调用方法

```
Type.NumMethod()

Type.Method()

Type.MethodByName('')

Method.Call(args)  

```
[code:-->DynamicMethod.go]()


---
### ✌ `Applications Of Reflection`

<p align="left">1、Go struct copy</p>
<p align="center">one - reflect</p>

```golang
func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}
```

---
<p align="center">two - json</p>

```golang
func main() {
	a := &A{1, "a", 1}
	//	b := &B{"b",2,2}

	aj, _ := json.Marshal(a)
	b := new(B)
	_ = json.Unmarshal(aj, b)

	fmt.Printf("%+v", b)
}
```
encode->json->decode

alternative plan: ffjson


---
### ✌ `Applications Of Reflection`

<p align="left">2、判断obj是否在target中</p>

```golang
// target支持的类型arrary,slice,map
func Contain(obj interface{}, target interface{}) (bool, error) {
    targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true, nil
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true, nil
        }
    }
    return false, errors.New("not in array")
}
```


---
### ✌ `Applications Of Reflection`

<p align="left">3、桥接模式<p>

<p align="left"><font size="5">不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数或方法反射。</font></p>

```golang

func bridge(funcPtr interface{}, args ...interface{})

```
<p align="left"><font size="5">funcPtr以接口的形式传入函数指针，函数参数args以可变参数的形式传入，bridge函数中可以用反射来动态执行funcPtr函数。</font></p>

[code-->bridge.go]()


---
### ✌ ️`Characteristics Of Reflection`

<font size="6">
<p align="left">1. Reflection can greatly improve the flexibility of the program. (interface{})</p>
<p align="left">2. Reflection uses the TypeOf and ValueOf functions to get the target object information from the interface.</p>
<p align="left">3. Reflection will treat anonymous fields as separate fields </p>

   [code-->AnonymousField.go]()
<p align="left">4. reflection to modify the state of the object, provided that interface.data is settable.(pointer-interface)</p>
<p align="left">5. The method can be called dynamically by reflection</p>
</font>


---
### ✌ `why reflect is slow`
<p align="center">java</p>

```java

Field field = clazz.getField("hello"); //java.lang.reflect.Field
field.get(obj1);
field.get(obj2);

```
<p align="left"><font size="5">取得的反射对象类型是可复用的。只要传入不同的obj，就可以取得这个obj上对应的 field。</font></p>

<p align="center">goland</p>

```go
type_ := reflect.TypeOf(obj)
field, _ := type_.FieldByName("hello") //reflect.StructField
```

```go
type_ := reflect.ValueOf(obj)  //malloc reflect.Value struct
fieldValue := type_.FieldByName("hello")
```


---
### ✌ `why reflect is slow`
[-->article link](http://legendtkl.com/2016/08/06/reflect-inside/)
 * <font color="#DAA520">Reflect Benchmark test</font>
 * <font color="#DAA520">Golang Profiling</font>

 * 涉及到内存分配以及后续的GC；
 * reflect实现里面有大量的枚举，也就是for循环，比如类型之类的。



---
### ✌ ️`Improve Reflection Performance`
[-->how to improve](https://studygolang.com/articles/12349)

<p align="left"><font color="#DAA520">Jsoniter</font>: 基于反射的 JSON 解析器,实现原理是用reflect.Type得出来的信息来直接做反射，而不依赖于reflect.ValueOf。</p>


---
* Jsoniter - <font color="#DAA520">struct</font>

```go

type TestObj struct {
	field1 string
}
struct_ := &TestObj{}
field, _ := reflect.TypeOf(struct_).Elem().FieldByName("field1") //reflect.StructField
field1Ptr := uintptr(unsafe.Pointer(struct_)) + field.Offset  //计算出字段的指针值
*((*string)(unsafe.Pointer(field1Ptr))) = "hello"
fmt.Println(struct_) 
//&{hello}

```


---
* Jsoniter - <font color="#DAA520">interface{}</font>

```go

type TestObj struct {
	field1 string
}
struct_ := &TestObj{}
structInter := (interface{})(struct_)
// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}
//从 interface{} 上取得结构体的指针
structPtr := (*emptyInterface)(unsafe.Pointer(&structInter)).word
field, _ := reflect.TypeOf(structInter).Elem().FieldByName("field1")
field1Ptr := uintptr(structPtr) + field.Offset  //计算出字段的指针值
*((*string)(unsafe.Pointer(field1Ptr))) = "hello"
fmt.Println(struct_)

```
* Jsoniter ? <font color="#DAA520">map</font>

[jsoniter文档](http://jsoniter.com/benchmark.html#optimization-used)

[代码地址](https://github.com/json-iterator/go)

---
### ✌ ️`Reference Link`

* [The Go Blog](https://blog.golang.org/laws-of-reflection)
* [Go语言的反射三定律](http://www.jb51.net/article/90021.htm)
* [官方reflect-Kind](https://golang.org/pkg/reflect/#Kind)
* [提高 golang 的反射性能](https://studygolang.com/articles/12349)

---
# Thanks


