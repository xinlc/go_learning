package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

// 协程比线程内存小, java jdk5以后，一个线程1M 协程只有2k
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go func() {
		// 	fmt.Println(i)  // 全部是10，变量会被共享, 只会是最后一个值
		// }()

		go func(i int) {
			fmt.Println(i)
		}(i) // 因为go是值传递，传进函数内的变量会被复制
	}
	time.Sleep(time.Millisecond * 50) // 避免测试程序结束的太快
}
