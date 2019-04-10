package unsafe_programming

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	// 不应该这样转
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	fmt.Println(unsafe.Pointer(&i))
	fmt.Println(f)
}

type MyInt int

// 合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	f := *(*[]MyInt)(unsafe.Pointer(&a))
	fmt.Println(f)
}

func TestAtomic(t *testing.T) {

	// 原子操作
	// atomic.StorePointer
}
