package map_ext_test

import (
	"fmt"
	"testing"
)

// Map val可以是函数
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

// 用Map实现Set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		fmt.Printf("%d is existing \n", n)
	} else {
		fmt.Printf("%d is not existing \n", n)
	}
	mySet[3] = true
	fmt.Println(len(mySet))

	delete(mySet, 1)
	n = 1
	if mySet[n] {
		fmt.Printf("%d is existing \n", n)
	} else {
		fmt.Printf("%d is not existing \n", n)
	}
	fmt.Println(len(mySet))
}
