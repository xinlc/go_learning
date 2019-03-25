package channel_close_test

import (
	"fmt"
	"testing"
	"time"
)

// 任务的取消

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} // {}{} 实例化一个空结构
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cannelled")
		}(i, cancelChan)
	}
	// cancel_1(cancelChan) // 只会取消一个
	cancel_2(cancelChan) // 全部都会取消
	time.Sleep(time.Second * 1)
}
