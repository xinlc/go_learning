package map_test

import (
	"fmt"
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	fmt.Println(m1[2])
	fmt.Println("len m1", len(m1))

	m2 := map[int]int{}
	fmt.Println("len m2", len(m2))

	m3 := make(map[int]int, 10)
	fmt.Println("len m3", len(m3))
}

// Map 在访问的Key不存在是，仍会返回零值，不能通过返回nil来判断是否存在
func TestAccessNotExitingKey(t *testing.T) {
	m1 := map[int]int{}
	fmt.Println(m1[1])
	m1[2] = 0
	fmt.Println(m1[2])

	// m1[3] = 0
	if v, ok := m1[3]; ok { // ok 返回true/false，代表key存不存在
		fmt.Println("key 3's value is ", v)
	} else {
		fmt.Println("Key 3 is not existing.")
	}
}

// Map遍历和数组一样
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
