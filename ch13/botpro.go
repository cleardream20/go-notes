package main

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void hello() {
    printf("hello, world!\n");
}

typedef struct {
    int x;
    int y;
} Point;

int sum(Point p) {
    return p.x + p.y;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// unsafeEx1()
	unsafeEx2()
    // C.hello() // 调用C函数
	// CStrEx()
	// structEx()
}

func unsafeEx1() {
	var i int
	var f float64
	var s string = "hello"
	
	fmt.Println("int size:", unsafe.Sizeof(i))      // 通常8字节(64位系统)
	fmt.Println("float64 size:", unsafe.Sizeof(f))  // 8字节
	fmt.Println("string size:", unsafe.Sizeof(s))   // 16字节(64位系统)
	
	type Person struct {
		Name string
		Age  int
	}
	fmt.Println("struct size:", unsafe.Sizeof(Person{}))  // 24字节(16+8)
}

func unsafeEx2() {
	type P struct {
		a bool
		b int16
		c []int
	}
	x := P{}
	fmt.Println("Sizeof:", unsafe.Sizeof(x.a), unsafe.Sizeof(x.b), unsafe.Sizeof(x.c))
	fmt.Println("Alignof:", unsafe.Alignof(x.a), unsafe.Alignof(x.b), unsafe.Alignof(x.c))
	fmt.Println("Offsetof:", unsafe.Offsetof(x.a), unsafe.Offsetof(x.b), unsafe.Offsetof(x.c))
}

func CStrEx() {
	// Go str 转 C
	goStr := "Hello, C!"
	cStr := C.CString(goStr)
	defer C.free(unsafe.Pointer(cStr)) // 必须手动释放内存

	// CString方法
	length := C.strlen(cStr)
	println("String length:", length)

	// C str 转 Go
	goStrBack := C.GoString(cStr)
	println(goStrBack)
}

func structEx() {
	p := C.Point{x: 10, y: 20}
    result := C.sum(p)
    fmt.Println("Sum:", result) // 输出: Sum: 30
}