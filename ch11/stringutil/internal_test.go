package stringutil

import "testing"

// 测试未导出的 reverseBytes 函数
func TestReverseBytes(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("a"), []byte("a")},
		{[]byte(""), []byte("")},
	}

	for _, tt := range tests {
		// 复制一份input，避免修改原数据
		inputCopy := make([]byte, len(tt.input))
		copy(inputCopy, tt.input)
		
		reverseBytes(inputCopy) // 直接调用未导出函数
		
		if string(inputCopy) != string(tt.want) {
			t.Errorf("reverseBytes(%q) = %q, want %q", 
				tt.input, inputCopy, tt.want)
		}
	}
}

// 同时也可以测试导出的Reverse函数
func TestReverse(t *testing.T) {
	if got := Reverse("hello"); got != "olleh" {
		t.Errorf("Reverse() = %q, want %q", got, "olleh")
	}
}