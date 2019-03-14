package loop_test

import (
	"fmt"
	"testing"
)

// Go做循环只有一个for关键字
// 不需要（）
func TestWhileLoop(t *testing.T) {
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// while(n<5)
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}

	// 无限循环 while(true)
	// n := 0
	// for {
	// 	fmt.Println(n)
	// }
}
