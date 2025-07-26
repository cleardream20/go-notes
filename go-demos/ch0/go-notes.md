# MY GO
golang酱! golang酱!! golang酱!!!😭

# 安装配置 go in vscode
## go 本体
[官网](https://golang.org/dl/)

环境变量应该会自动设置
1. 用户变量 GOPATH  e.g. C:/Users/xxx/go
2. 系统变量 PATH  e.g. xxx/Go/bin

验证安装：

命令行里
```bash
go version
```

## go plugins in vscode
插件商店搜 go，下载安装就OK

启动？
B站视频：
[Getting started with VS Code Go](https://www.bilibili.com/video/BV11x4y1p7Fw?vd_source=127961da3fc6c308c415223bd57e3f44)

依赖下载失败？
网络原因（科学上网好像也不太行，不知道为什么，不过可以再逝逝），需要设置一下国内代理

vscode 打开 powershell
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

详情见此处哦喵~
[国内代理~](https://goproxy.cn/)

还是不行，那就叉掉vscode重进 / 重下 / 一个一个下，不要一下全都下了...
（反正一顿操作之后总算下载完了）

## 大mod盛世
go.mod改变世界？🥳

```bash
# 进入相应文件夹/工作区/项目根目录
cd ...(替换为你的go文件夹)
# go mod init命令创建
go mod init xxx(.com)/xxx/xxx
# e.g. go mod init example.com/hajimi/manbo
```

如果生成相应的go.mod文件，说明成功！
点进去你的go.mod，看看长啥样的~

示例go.mod文件
```go
module github.com/your/project  // 模块路径（必须）

go 1.21                        // 最低Go版本要求

require (
    github.com/pkg/errors v0.9.1 // 直接依赖
    golang.org/x/sync v0.3.0
)

replace (
    // 本地替换（调试时常用）
    golang.org/x/net => ./local/net
)

exclude (
    // 排除特定版本
    github.com/old/lib v1.2.3
)
```

## Code Runner
vscode 插件市场搜 code runner
34.5M下载量 你值得拥有！

如名，啥都能跑
右键，第一行Run Code(Ctrl+Alt+N)

# go 教学
## [《Go语言圣经》](https://golang-china.github.io/gopl-zh/)

[《The Go Programming Language》](https://www.gopl.io/)

## 梦开始的地方——Hello World
```go
package main

import "fmt"

func main() {
    /* 注释是C-Type的！ */
    fmt.Println("Hello, World!") // 回家了孩子们
}
```

啧啧啧，21世纪C语言😘

## 运行
1.Code Runner 现学现卖

---
2.命令行 run
```bash
go run hello.go
```

```bash
go run .
```
（一般涉及到多个文件联动时）

---
3.命令行 build(类C++)
```bash
go build hello.go
# 会生成一个hello.exe
./hello.exe
```

... more info in 圣经 chapter10 包和工具 10.7工具

<!-- fitten code 插件推荐？ -->

<!-- 不如cursor?... -->

## Chap1 入门
### 1.1 Hello, Wolrd?
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

（下面的很多内容会在ch10 package&tools里面详细说，敬请期待~😘）
#### package
Go语言代码通过包(package)组织，类似其他语言的库(libraries)或者模块(modules)
一个package由单个目录下一个/多个.go文件组成
每个.go抬手都是一句`package ...`，表示文件属于哪个包

`main`包（面包？😋）比较特殊。它定义了一个独立可执行的程序，而不是一个库。
`main`里面的`main`函数——程序执行时的入口（类C！🥰）

#### import
后面再跟`import`，即C`#include` or Python`import`
注意Go高度严格的代码格式要求，`import`必须在`package`后面
而且，`import`多了或少了，程序都无法编译通过！（万能头一败涂地😇）

为什么严格？严格 = 标准 = 统一 = 少争议，你的free会让你issue乱飞😋
极致的标准🤔 = 极致的轮椅🤤

#### 函数
函数声明
func关键字、函数名、参数列表、返回值列表、{函数体}

#### 句末分号？
不需要添~~绝类python  都没有=都有

真不需要？实际上，编译器会主动把特定符号后的换行符转换为分号

一行有多条语句？老实加上

more info？下...回分解！😋

#### 代码格式化？
`gofmt`工具，很多文本编辑器都可以配置为保存文件时自动执行`gofmt`，这样你的源代码总会被恰当地格式化

ctrl-s保存自动代码格式化？来逝逝吧！

---
`goimports`工具，可以根据代码需要，自动地添加或删除 import 声明

```bash
go get golang.org/x/tools/cmd/goimports
```

more go tools? 下下下...回分解！😋😋

