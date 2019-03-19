package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// string 常用函数
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")

	for _, part := range parts {
		fmt.Println(part)
	}

	fmt.Println(strings.Join(parts, "-"))

}

func TestCov(t *testing.T) {
	s := strconv.Itoa(10)
	fmt.Println("str" + s)

	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10 + i)
	}
}
