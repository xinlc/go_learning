package condition_test

import (
	"fmt"
	"testing"
)

// if和其他语言一样
// Go的if支持两段写法，赋值和比较 ,多用于方法多返回值
func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		fmt.Println("1==1")
	}

	// 方法多返回值并赋值判断
	// if v, err := someFun(); err == nil {

	// } else {

	// }
}

// switch 不限制常量和整数, 可以用string
// switch 默认有break
// case 后支持多个项
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			fmt.Println("Even")
		case 1, 3:
			fmt.Println("Odd")
		default:
			fmt.Println("it is not 0-3")
		}
	}
}

// switch可以当if用, 可以写表达式
func TestSwitchCaseConditon(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			fmt.Println("Even")
		case i%2 == 1:
			fmt.Println("Odd")
		default:
			fmt.Println("unknow")
		}
	}
}
