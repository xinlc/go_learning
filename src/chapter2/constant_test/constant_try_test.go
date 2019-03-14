package constant_test

import (
	"fmt"
	"testing"
)

// 快速设置连续值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
	fmt.Println(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 // 0111
	// a := 1 // 0001
	fmt.Println(a&Readable == Readable, a&Writable == Writable)
}
