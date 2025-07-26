package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"example.com/go-demo/ch12/utils"
)

func main() {
	// refTandV()
	// val2inter()
	// DisplayEx()
	// modifyValue()
	Print(os.Stdout)
}

func refTandV() {
	var w io.Writer = os.Stdout
	t := reflect.TypeOf(w) // 动态值
	v := reflect.ValueOf(w) // 动态类型
	fmt.Println(t, v)
	// vs fmt.Printf
	fmt.Printf("%T %v\n", w, w) // %T Type, %v value
	// 其实就是靠reflect.T/V实现的

	// 我们再来看看t和v都是什么类型的
	tt := reflect.TypeOf(t) // *reflect.rtype (reflect.Type)
	tv := reflect.TypeOf(v) // reflect.Value
	fmt.Println(tt, tv)

	// reflect.Type和reflect.Value类型也实现了String()方法（满足fmt.Stringer接口）
	fmt.Println(t.String(), v.String()) // *os.File <*os.File Value>

	// reflect.Value.Type() = reflect.Type
	fmt.Println(v.Type(), t) // 一样的，*os.File *os.File
}

func val2inter() {
	// interface->value, then value->interface?
	var num int = 3 // <int, 3>
	v := reflect.ValueOf(num) // reflect.Value
	x := v.Interface() // interface{}
	i := x.(int) // int
	fmt.Println(i)
}

func DisplayEx() {
	strangelove := utils.Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	utils.Display("strangelove", strangelove)
}

func modifyValue() {
    var x float64 = 3.4
    v := reflect.ValueOf(&x).Elem() // 必须获取可设置的Value
    
    if v.CanSet() { // CanSet()检查是否是可取地址并可被修改的
        // v.Set(reflect.ValueOf(7.1))
        v.SetFloat(7.1)
        // 这两种set方法都OK！
        // 还有SetInt(), SetString()...
        fmt.Println(x) // 输出: 7.1
    }
}

func Print(x interface{}) {
    v := reflect.ValueOf(x)
    t := v.Type()
    fmt.Printf("type %s\n", t)

    for i := 0; i < v.NumMethod(); i++ {
        methType := v.Method(i).Type()
        fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
            strings.TrimPrefix(methType.String(), "func"))
    }
}

