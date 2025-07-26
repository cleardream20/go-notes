package main

import "fmt"

func main() {
	deferTest()
	deferStack()
	valInDefer()

	// panicTest()
	panicRecover()
	fmt.Println("你不执行我执行")
	outer()
	fmt.Println("你说我能不能执行？")
}

func deferTest() {
	defer fmt.Println("这是最后执行的语句")

	fmt.Println("这是第一个执行的语句")
	fmt.Println("这是第二个执行的语句")
}

func deferStack() {
	defer fmt.Println("第一个 defer")
	defer fmt.Println("第二个 defer")
	defer fmt.Println("第三个 defer")
}

func valInDefer() {
	i := 1
	defer fmt.Println("i in Defer:", i)
	i++
	fmt.Println("i in Normal:", i)
}

func panicTest() {
	panic("发生了严重错误！") // 程序在此处中断
	println("这行不会执行") // unreachable code
}

func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: ", r)
		}
	}() // lambda函数？！

	panicTest()
	fmt.Println("这行不会执行")
}

func inner() {
	defer fmt.Println("inner defer")
	panic("inner panic")
}

func outer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("outer recovered:", r)
		}
	}()
	inner()
}
