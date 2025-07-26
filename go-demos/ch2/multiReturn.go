package main

import "errors"

func swap(a, b int) (int, int) {
	return b, a
}

func divide(a, b float64) (float64, error) {
	if b==0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil // nil类似于null
}

func calc(a, b int) (sum int, diff int, prod int) {
	sum = a + b
	diff = a - b
	prod = a * b
	return //自动返回已命名的变量
}

func getData() (int, string, bool) {
	return 1, "str", true
}