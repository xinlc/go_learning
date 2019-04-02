package util_anyone_reply_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 仅需任意任务完成
// 如：用户同时向Baidu、Google、Bing发起搜索请求，只要有一个返回就可以了

func runTask(id int) string {
	// fmt.Println(id)
	time.Sleep(10 * time.Microsecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// ch := make(chan string)   // 可能会导致协程序泄露
	ch := make(chan string, numOfRunner) // 防止协程泄露
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch // 只要chan有数据 就返回, 注意：会有协程泄露，因为只有一个chan数据被取走，其他的没人取走协程被阻塞, 采用buffe chan可以防止协程序泄露
}

func TestFirstResponse(t *testing.T) {
	fmt.Println("Before:", runtime.NumGoroutine()) // 打出协程数
	fmt.Println(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After:", runtime.NumGoroutine())
}
