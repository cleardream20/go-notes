# cobra-cli

[B站教程（轮椅）](https://www.bilibili.com/video/BV1Qi2eYuEyX?vd_source=127961da3fc6c308c415223bd57e3f44)

（视频是从YouTube上搬运的）
[MightyMoud老哥YouTube主页](https://www.youtube.com/@MightyMoud)

# 启动
get
```sh
go install github.com/spf13/cobra-cli@latest
```

mkdir and cd
```sh
mkdir go-cli
cd go-cli
```

go mod init
```sh
go mod init example.com/go-cli
```

cobra-cli init
```sh
cobra-cli init
```

`go-cli/cmd/root.go` is to be operated

# 运行
run main.go
```sh
go run main.go
```

local flags
```go
rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

Run: func(cmd *cobra.Command, args []string) {
		toggle, _ := cmd.Flags().GetBool("toggle")
		if toggle {
			fmt.Println("Toggle is true")
		}
		fmt.Println("Root Called")
	},
```

```sh
go run main.go --toggle

go run main.go -t
```

# add命令
add parent
```sh
cobra-cli add parent
```

新建cmd/parent.go文件

像修改root.go一样修改parent.go

运行
```sh
go run main.go parent

go run main.go parent --toggle

go run main.go paren -t
```

# 命令 & 子命令
```sh
docker compose up
```
`docker` CLI的名称
`compose` 命令名称
`up` 子命令名称
层次结构？！

init() in `parent.go`
```go
// 将parent命令添加到了root命令
rootCmd.AddCommand(parentCmd)
// 可以如法炮制，将子命令添加到父命令
// 不过也不用，使用cobra-cli工具
```

cobra 添加子命令
```sh
cobra-cli add child -p parentCmd
```

# Pterm
github.com/pterm/pterm
```go
go get github.com/pterm/pterm
```
接收用户输入，显示表格

# gorm
gorm.io/gorm
```sh
go get gorm.io/gorm
```

# glebarez/sqlite
```sh
go get github.com/glebarez/sqlite
```

# sqlite
[sqlite官网下载](https://www.sqlite.org/download.html)

# Cobra的命令架构？
## main structure
Commands, Args and Flags
```sh
APPNAME COMMAND ARG --FLAG

APPNAME VERB NOUN --ADJECTIVE
```

e.g.
```go
hugo server --port=1313
```
`hugo`:name, `server`:command, `port`:flag

```go
git clone URL --bare
```
`git`:name, `clone`:command, `URL`:arg, `bare`:flag

## Commands
central point

在`xx/cmd`文件夹下，cmd?🤔 command!😋

所有交互——包含在命令中

命令可以有子命令

命名：驼峰命名法`addUser` √
snake_case`add_user`，kebab-case`add-user`×

## 

# License & Copyright?
a custom license?
```go
author: Steve Francia <spf@spf13.com>
year: 2020
license:
  header: This file is part of CLI application foo.
  text: |
    {{ .copyright }}

    This is my license. There are many like it, but this one is mine.
    My license is my best friend. It is my life. I must master it as I must
    master my life.
```

claim Copyright ©
```go
Copyright © 2020 Steve Francia <spf@spf13.com>

This is my license. There are many like it, but this one is mine.
My license is my best friend. It is my life. I must master it as I must
master my life.
```

in the header of files...
```go
/*
Copyright © 2020 Steve Francia <spf@spf13.com>
This file is part of CLI application foo.
*/
```

# 像docker一样起名字——mycmd
在 Windows 系统上实现类似 `docker` 这样的全局命令效果，需要以下 **4 个步骤**：

（省流：1.`go build -o mygo.exe`生成可执行文件，2.把`mygo.exe`所在路径加入到系统环境变量`PATH`中去，就可以随时随地执行`mygo`啦~）

---

### **🛠️ 操作步骤（Windows 专属版）**
#### **1. 编译生成 exe 文件**
在项目目录下打开 PowerShell 或 CMD，执行：
```powershell
go build -o mycmd.exe  # 会生成 mycmd.exe 可执行文件
```
（将 `mycmd` 替换为你想要的命令名，比如 `aihelper.exe`）

#### **2. 临时测试运行**
```powershell
.\mycmd.exe --toggle  # 测试是否能正常运行
```

#### **3. 添加到系统 PATH（实现全局命令）**
##### **方法一：快速配置（当前用户有效）**
1. 右键点击 **此电脑** → **属性** → **高级系统设置** → **环境变量**
2. 在 **用户变量** 中找到 `PATH`，点击 **编辑**
3. 点击 **新建**，输入你的 exe 文件所在目录的**完整路径**（如 `C:\Users\你的用户名\go-cli`）
4. 一路点击 **确定** 保存

##### **方法二：专业推荐（所有用户有效）**
1. 将 `mycmd.exe` 复制到系统程序目录：
   ```powershell
   # 以管理员身份运行 PowerShell
   Copy-Item .\mycmd.exe "C:\Windows\System32\"
   ```
   之后所有用户都可以直接使用 `mycmd` 命令。

#### **4. 验证效果**
**重新打开一个新的 CMD/PowerShell 窗口**（重要！），然后直接运行：
```powershell
mycmd -h      # 查看帮助
mycmd --toggle  # 测试功能
```
如果看到输出，说明配置成功！

---

### **💡 实用技巧**
#### **1. 修改命令显示名称**
编辑 `root.go`，将 `Use: "go-cli"` 改为你的命令名：
```go
Use:   "mycmd",  // 用户输入的命令名
Short: "我的专属命令行工具",  // 简短描述
```

#### **2. 一键编译并安装**
创建 `install.ps1` 脚本（PowerShell）：
```powershell
go build -o mycmd.exe
Copy-Item .\mycmd.exe "C:\Windows\System32\" -Force
Write-Host "安装成功！现在可以全局使用 mycmd 命令了" -ForegroundColor Green
```
右键选择 **使用 PowerShell 运行** 即可自动完成安装。

#### **3. 卸载方法**
删除系统目录中的 exe 文件：
```powershell
# 管理员身份运行
Remove-Item "C:\Windows\System32\mycmd.exe"
```

---

### **🚨 注意事项**
- **杀毒软件警告**：首次运行时可能会被 Windows Defender 拦截，选择 **允许运行** 即可。
- **命令冲突**：确保你的命令名（如 `mycmd`）不和系统已有命令（如 `ping`）重复。
- **修改代码后**：需要重新执行 `go build` 生成新的 exe 文件。

---

现在你的 Go 程序已经和 `docker`/`git` 一样成为系统级命令了！试试在任意目录打开终端输入 `mycmd` 吧！ 🎉




