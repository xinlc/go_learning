package csp_test

import (
	"fmt"
	"testing"
	"time"
)

// Communicating Sequential Processes
// channel实现异步任务

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is doen.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// Go channel 有两种
// 1. 放一个取一个
// 2, 有个buffer 可以放多个
func AsyncService() chan string {
	retCh := make(chan string) // 会阻塞
	// retCh := make(chan string, 1) // buffer channel, 没达到cap数 不会阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret // 往channel放数据
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 从channel取数据
	time.Sleep(time.Second * 1)
}
