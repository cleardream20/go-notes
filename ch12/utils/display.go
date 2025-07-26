package utils

import (
	"fmt"
	"reflect"
)

func Display(path string, vi interface{}) {
	v := reflect.ValueOf(vi)
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	// 对Slice和Array
	// 递归处理每一个元素
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			Display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	// 对于Struct
	// 递归处理每一个成员
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			Display(fieldPath, v.Field(i))
		}
	// 对于指针
	// 打印时多一个(*...)
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			Display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	// 对于接口
	// 打印动态类型，递归处理动态值
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			Display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %v\n", path, v.Interface())
	}
}

type Movie struct {
    Title, Subtitle string
    Year            int
    Color           bool
    Oscars          []string
    Sequel          *string
}