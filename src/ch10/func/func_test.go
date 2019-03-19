package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 函数在Go中是一等公民
// 函数传值都是值传递
// 可以多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFnc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	fmt.Println(a, b)

	tsSF := timeSpent(slowFnc)
	fmt.Println(tsSF(10))
}

// 可以边长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	s := Sum(1, 2, 3, 4, 5)
	fmt.Println(s)
}

func Clear() {
	fmt.Println("Clear resources.")
}

// 延迟执行函数
// 类似java中finally, 可以做释放资源等等
func TestDefer(t *testing.T) {
	// defer Clear()
	defer func() { // 匿名函数
		fmt.Println("Clear resources.")
	}()
	fmt.Println("Start")
	panic("err")
}
