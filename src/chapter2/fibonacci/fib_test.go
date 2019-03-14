package fib

import "testing"
import "fmt"

func TestFirstTry(t *testing.T) {
	// 以下都是赋值写法
	var a int = 1
	var b int = 1
	// var (
	// 	a int = 1
	// 	b int = 1
	// )
	// a := 1
	// b := 2
	// t.Log(a)
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
		// a, b = b, a  // 变量值交换
	}
	fmt.Println()
}
