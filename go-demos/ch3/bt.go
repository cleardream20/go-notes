package main

import "fmt"

func main() {
	printFloat()
}

func printFloat() {
	// var f float64 = 114.514
	f := 114.514
	fmt.Println("====基础print====")
	fmt.Println(f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%v\n", f)
	fmt.Printf("%#v\n", f)

	// 控制小数位数
	fmt.Println("====控制小数位数====")
	fmt.Printf("%.2f\n", f) // 两位小数，自动四舍五入
	// 但是四舍五入稍微有点问题（估计是二进制的问题）
	// 可以尝试分别打印114.514, 114.515, 114.516

	// 控制总宽度
	fmt.Println("====控制宽度&对齐====")
	fmt.Printf("%10.3f\n", f) // 右对齐
	fmt.Printf("%-10.3f\n", f) // 左对齐
	fmt.Printf("%010.3f\n", f) // 填充前导0

	// 正负号？
	fmt.Println("====控制正负号====")
	fmt.Printf("%+f\n", f)
	fmt.Printf("%-f\n", -f) // 这个纯在玩了
	fmt.Printf("%+f\n", -f) // 想蒙混过关？该罚！

	// 科学计数法
	fmt.Println("====科学计数法====")
	fmt.Printf("%e\n", f)
	fmt.Printf("%.3e\n", f) // 保留3位小数

	// 自动选择格式
	fmt.Print("====","%","g","自动选择是否使用科学计数法====\n")
	fmt.Printf("%g\n", 123456.789) // 自动选择 %f 或 %e: 123456.789
	fmt.Printf("%g\n", 1.23456789e5)
	fmt.Printf("%g\n", 1.23456789e8) // 输出: 1.23456789e+08
	// 短的就不用，长的就用
}