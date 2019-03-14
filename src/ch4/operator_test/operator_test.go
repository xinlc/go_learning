package operator_test

import (
	"fmt"
	"testing"
)

// 1. Go语言中不支持 前置++，-- e.g. ++a, --a
// 2. 数组==比较，同为维度且含有相同个数元素的数组才能比较
// 3. 数组每个元素都相等才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 5}
	fmt.Println(a == b)
	fmt.Println(a == c)
}
