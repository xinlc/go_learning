package try_test

import "testing"

// 1. 测试文件名已_test.go结尾
// 2. 测试方法已Test开头，第一个字母大写代表可以包外访问
func TestFirstTry(t *testing.T) {
	t.Log("My first try")
}
