package main

import "fmt"

func modifyArray(arr [3]int) { // 值传递
    arr[0] = 100 // 只修改副本
	// unused write to array index 0:int
}

func modifySlice(s []int) { // 引用类型
	s[0] = 100
}

func lensTypes() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b) // true 长度相等值也相等
	// c := a + b  ×，不能直接加

	// c := [5]int{1, 2, 3, 4, 5}
	// d := [3]int{1, 2, 3}
	// fmt.Println(c == d) // × invalid operation: c == d (mismatched types [5]int and [3]int)
	// 长度不等直接不能 == 比较

	b = [3]int{4, 5, 6}
	fmt.Println(a == b) // false 长度相等，值不相等也不相等
}

func main() {
	a := [3]int{1, 2, 3}
    modifyArray(a)
    fmt.Println(a) // 输出: [1 2 3]（原数组未改变）

	lensTypes()

	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println(s)
}