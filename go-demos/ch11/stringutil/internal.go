package stringutil

import "fmt"

// 未导出的内部函数（首字母小写）
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// 导出的公共函数
func Reverse(s string) string {
	b := []byte(s)
	reverseBytes(b) // 调用内部实现
	return string(b)
}

func ToDigit(s string) (int, error) {
    if len(s) == 0 {
        return 0, fmt.Errorf("empty string")
    }

    var num int = 0
    for i := 0; i < len(s); i++ {
        if s[i] < '0' || s[i] > '9' {
            return 0, fmt.Errorf("invalid character '%c'", s[i])
        }
        num = num*10 + int(s[i]-'0')
    }
    return num, nil
}