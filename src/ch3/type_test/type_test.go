package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64 // 定义类型别名

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a   Go语言不支持隐身类型转换
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	fmt.Println(a, b, c)
}

// Go 支持自动垃圾回收，也支持指针，但不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 不支持指针运算
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
	fmt.Println()
}

// Go 中string类型默认值是空字符串 不是nil
func TestString(t *testing.T) {
	var s string
	fmt.Println("*" + s + "*")
	if s == "" {
		fmt.Println(len(s))
	}

}
