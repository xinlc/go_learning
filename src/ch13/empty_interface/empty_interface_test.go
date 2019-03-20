package empty_interface_test

import (
	"fmt"
	"testing"
)

// 空接口可以代表任何类型 类似java中的Object
// 通过断言类将空接口转换为指定类型
// v, ok := p.(int) // ok=true 时为转换成功
func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok { // 断言
	// 	fmt.Println("Interface", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("String", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")

	// 简化写法
	switch v := p.(type) {
	case int:
		fmt.Println("Interface", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyIntefaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

// Go 接口最佳实践
// 1. 倾向于使用小的接口定义，很多接口只包含一个方法
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Write interface {
	Write(p []byte) (n int, err error)
}

// 2. 较大的接口定义，可以由多个小接口定义组合而成
type ReadWriter interface {
	Reader
	Write
}

// 3. 只依赖于必要功能的最小接口
// func StoreDate(reader Reader) error {
// }
