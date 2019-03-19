package array_test

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5} // 自己计算长度
	// arr4 := [2][2]int{{1, 2}, {3, 4}} // 多维数组
	fmt.Println(arr[1], arr1[2], arr3[3])
	// fmt.Println(arr[1], arr1[2], arr3[3])
}

// Go语言中声明的变量必须使用，否则编辑不通过，但可以用 _ 占位，表示我们不关系这个
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for idx, e := range arr3 {
		fmt.Println(idx, e)
	}

	for _, e := range arr3 {
		fmt.Println(e)
	}
}

// 数组截取
// a[开始索引(包含), 结束索引(不包含)]
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:3]     // 取前三个, [3:]第三个后的所有元素, 不支持负数
	arr := arr3[1:len(arr3)] // 2,3,4,5
	fmt.Println(arr3_sec, arr)
}
