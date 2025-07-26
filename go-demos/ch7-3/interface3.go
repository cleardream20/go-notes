package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// typeAssertEx()
	// nil2err()
	IfStringerEx()
	processInputEx()
}

func typeAssertEx() {
	var w io.Writer
	w = os.Stdout // *os.File类型
	f, ok := w.(*os.File)
	fmt.Printf("%#v, %#v\n", f, ok)
	c, ok := w.(*bytes.Buffer)
	fmt.Printf("%#v, %#v\n", c, ok)
}

func nil2err() {
	var w io.Writer  // 声明但未赋值的接口变量，此时是nil接口值
	
	// 尝试对nil接口值进行类型断言
	f, ok := w.(*os.File)
	fmt.Printf("断言*os.File: (%#v, %v)\n", f, ok)  // (nil, false)

	// 甚至断言为interface{}也会失败
	a, ok := w.(interface{})
	fmt.Printf("断言interface{}: (%#v, %v)\n", a, ok)  // (nil, false)

	w = os.Stdout
	// 对非nil接口值进行类型断言
	f, ok = w.(*os.File)
	fmt.Printf("断言*os.File: (%#v, %v)\n", f, ok)  // (nil, false)
}

type Stringer interface {
    String() string
}

// 检查v是否实现Stringer接口
func printIfStringer(v interface{}) {
    if s, ok := v.(Stringer); ok {
        fmt.Println(s.String())
    } else {
        fmt.Printf("%v 没有实现 Stringer 接口\n", v)
    }
}

type S1 struct {
	Str string
}

func (s S1) String() string {
	return s.Str
}

type S2 struct {
	Str string
}

// func (s S2) String() string {
// 	return s.Str
// }

func IfStringerEx() {
	var s1 S1 = S1{Str: "123"}
	printIfStringer(s1)
	var s2 S2
	printIfStringer(s2)
}

func processInput(input interface{}) {
    switch v := input.(type) {
    case int:
        fmt.Printf("整数: %d\n", v*2)
    case string:
        fmt.Printf("字符串长度: %d\n", len(v))
    case bool:
        fmt.Printf("布尔值: %t\n", v)
    default:
        fmt.Printf("不支持的类型: %T\n", v)
    }
}

func processInputEx() {
	processInput(1)
	processInput("2")
	processInput(true)
	processInput(3.14)
	processInput(S1{Str: "123"})
	// 甚至nil也可以传！
	processInput(nil)
}