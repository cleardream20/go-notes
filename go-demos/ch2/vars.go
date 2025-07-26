package main

import "fmt"
import "unsafe"
import "errors"

import pkg "example.com/go-demo/ch2/pkg"

func main() {
	var a = 1
	const b = 2
	var c int = a*a-b
	d := a+b
	fmt.Printf("vars: a = %d, b = %d, c = %d, d = %d\n", a, b, c, d)
	fmt.Printf("square a = %d\n", square(a));

	// 变量名必须以Unicode字母或下划线开头
	var 姓名 string = "未知"
	fmt.Printf("姓名：%s\n", 姓名)

	点兵()

	// 获取另一个main包下的consts.go中的数据
	fmt.Println(PublicConst)
	fmt.Println(privateConst)

	declarations()

	mulReEx()

	pkgImport()
}

func square(a int) int {
	return a*a
}

func 点兵() {
	fmt.Println("====变量12卷, 卷卷有var名====")
	/* 整型 */
	var a int = 1; var b int8 = 1; var c int16 = 1;
	var d int32 = 1; var e int64 = 1;
	// var l int128 = 1;  ×没有int128!
	// print(a, b, c, d, e);
	fmt.Println(a, b, c, d, e)

	/* 无符号整型 */
	var ua uint = 1; var ub uint8 = 1; var uc uint16 = 1;
	var ud uint32 = 1; var ue uint64 = 1;
	fmt.Println(ua, ub, uc, ud, ue);

	// uintptr Go底层编程工具
	// 存储指针的整数值，需谨慎使用
	// 存储指针值，但不是指针！GC(Gabage Colleciton)会忽略它！
	x := 1
	p := &x
	ptr := uintptr(unsafe.Pointer(p))
	// 指针 -> unsafe.Pointer -> uintptr
	// println(ptr)
	fmt.Printf("x地址: 0x%x\n", ptr)

	/* 浮点型 */
	var fa float32 = 1.2; var fb float64 = 1.2;
	fmt.Println(fa, fb);

	/* 布尔型 */
	var yes bool = true; var no bool = false;
	fmt.Println(yes, no)

	/* 复数complex型 */
	// complex64 = float32实部 + float32虚部
	// complex128 = float64实部 + float64虚部
	var c1 complex64 = complex(3, 4)
	var c2 complex128 = complex(1.5, 2.7)
	fmt.Println(c1, c2)
	// real实部 imag虚部
	fmt.Println(real(c2), imag(c2))

	// byte 字节 = int8
	// 常用于处理二进制数据 / ASCII字符
	var ch byte = 'A';
	fmt.Println(ch)

	// rune Unicode码点 = int32
	// 用于处理UTF-8编码的字符串
	// 任何Unicode字符，包括中文、emoji！（大win特win）
	var r rune = '中'
	fmt.Println(r)
	var emoji = '😋'
	fmt.Println("emoji:", emoji)

	// errors 有点像C/C++里的Exception
	var err error = errors.New("报个错")
	fmt.Println(err)

}

func multiReturn(a int) (int, int) {
	return a+1, a+2
}

func declarations() {
	// var name type = value
	fmt.Println("====该选择怎样的声明呢？====")
	var s string
	fmt.Println(s) // ""
	var i, j, k int
	var b, f, str = true, 2.3, "four" // bool, float64, string
	fmt.Println(i, j, k, b, f, str);

	var a = 1
	// 多返回值，类python
	_a, __a := multiReturn(a)
	fmt.Println(a, _a, __a)


}

func mulReEx() {
	fmt.Println("====我们需要更多的返回值====")
	// funcs in file 'multiReturn.go'
	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)
	// := 变量声明  = 赋值操作
	a, b = b, a
	fmt.Println(a, b)

	c, err := divide(1, 0)
	fmt.Println(c, err)

	s, d, p := calc(2, 3)
	fmt.Println(s, d, p)

	// 不需要的用 _ 占个位
	_, str, _ := getData()
	print(str)
}

func pkgImport() {
	fmt.Println("====本地包导入示例====")
	fmt.Println(pkg.Add(1, 2))
	// fmt.Println(pkg.prod(2, 3)) × undefined: pkg.prod
	// 大小写暗含公有/私有  能否被其他包使用/访问
	// 大写Public  小写pivate
}