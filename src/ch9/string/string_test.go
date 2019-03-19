package string_test

import (
	"fmt"
	"testing"
)

// string 是不可变的byte slice
// rune 是能取出字符串里的unicode
// Unicode 是编码集, UTF-8 是物理存储实现
func TestString(t *testing.T) {
	var s string
	s = "hello"
	fmt.Println(s)
	fmt.Println(len(s))
	// s[1] = '3' // string是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	fmt.Println(s)
	fmt.Println(len(s))

	s = "中"
	c := []rune(s) // rune 字符->unicode切片
	fmt.Printf("中 unicode %x", c[0])
	fmt.Println()
	fmt.Printf("中 UTF8 %x", s)
	fmt.Println()
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c %[1]x", c)
		fmt.Printf("\t %[1]d", c)
		fmt.Println()
	}
}
