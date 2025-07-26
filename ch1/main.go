package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	utils "example.com/go-demo/ch1/utils"
)

func main() {
	// cmdArgs()
	// cmdArgsFlag()
	// theSameLines()
	// utils.MyGif()
	// utils.GetUrl()
	// utils.GetUrlConcur()
	utils.WebServeEx()
}

func cmdArgs() {
	for i, arg := range os.Args {
		fmt.Printf("参数 %d: %s\n", i, arg)
	}
}

func cmdArgsFlag() {
    // 定义参数
    name := flag.String("name", "Guest", "用户名")
    age := flag.Int("age", 0, "用户年龄")
    verbose := flag.Bool("v", false, "详细模式")
    
    // 解析参数
    flag.Parse()
    
    fmt.Printf("姓名: %s, 年龄: %d, 详细模式: %v\n", *name, *age, *verbose)
    fmt.Println("剩余参数:", flag.Args())
}

func theSameLines() {
	// 统计数据
    counts := make(map[string]int)
	// 提供命令行参数获取文件路径
    for _, filename := range os.Args[1:] {
		// os.ReadFile()打开文件
        data, err := os.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
		// 过滤空行，分割成一行一行的数据
        for _, line := range strings.Fields(string(data)) {
   			counts[line]++
		}
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
