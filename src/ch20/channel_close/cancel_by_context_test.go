package channel_close_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 关联任务的取消
// Context 在go 1.9版本后加入

func isCancelled2(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()) // Background 获取根节点
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled2(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cannelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
