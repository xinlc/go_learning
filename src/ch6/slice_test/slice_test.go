package slice_test

import (
	"fmt"
	"testing"
)

// 切片和数组很类似
// 数组VS切片：
// 1.切片容量可伸缩
// 2.切片不能进行比较
// 切片结构：元素指针（指向一片可以变长数组），len，cap
// len : 元素个数 , cap: 容量个数（有可能没初始化，是不能访问的）
func TestSliceInit(t *testing.T) {
	// 切片的声明
	var s0 []int
	fmt.Println(len(s0), cap(s0))
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s1), cap(s1))

	s2 := make([]int, 3, 5) // 类型，len，cap
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2[0], s2[1], s2[2])
}

// 切片自动增长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i) // 记得要赋值给s,会创建新的存储空间并自动复制
		fmt.Println(len(s), cap(s))
	}
}

// 共享空间
func TestSliceShareMemory(t *testing.T) {
	year := []string{"jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer))

	// 可以看到修改summer内容，Q2和year都发生变化，因为他们使用的是同一块空间
	summer[0] = "Unknow"
	fmt.Println(Q2)
	fmt.Println(year)
}

func TestSliceCompare(t *testing.T) {
	// 切片不能比较
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// if a == b {

	// }
}
