package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    var a, b, c int = 1, 2, 3
    fmt.Println(a, b, c)

    a, d := 2, 4
    fmt.Println(a, d)

    s := "hello, wolrd!"
    for i, r := range s {
        fmt.Printf("%d\t%q\t%d\n", i, r, r) // 下标，UTF-8字符，UTF-8字符对应的值
    }
}