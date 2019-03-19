package interface_test

import (
	"fmt"
	"testing"
)

// 定义接口
// Go的实现接口并没有用implement，而是用了duck type
// Go接口 VS Java接口
// 1. 接口为非入侵性，实现不依赖接口定义
// 2. 所有接口的定义可以包含在接口使用者包内
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// duck type 方法签名一样就可以了
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hellow World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
}
