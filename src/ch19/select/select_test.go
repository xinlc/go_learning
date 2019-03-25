package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

// 多渠道选择，实现超时
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		fmt.Println(ret)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("time out")
		// default:
		// 	fmt.Println("default")
	}

}
