package main

import "fmt"
import "os"
import "image"
import _ "image/png"  // 添加 PNG 解码支持
import _ "image/jpeg" // 添加 JPEG 解码支持

// 不需要/png /jpeg里面具体的函数/变量什么的
// 只需要让image知道有PNG JPEG就OK

func main() {
    f, _ := os.Open("eg.jpg")
    // img, format, _ := image.Decode(f) // 自动识别格式
    _, format, _ := image.Decode(f)
	// fmt.Println(img)
	fmt.Println("图像格式:", format)
}