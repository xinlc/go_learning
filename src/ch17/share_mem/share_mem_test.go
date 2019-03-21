package share_mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Go中的并发控制和其他语言一样，利用锁机制
// sync.Mutex    互斥锁
// sync.RWLock   对读锁和写锁分开, 比Mutex性能高一些
// sync.WaitGroup 要等待的任务

func TestCount(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)            // 等待所有协程执行完，避免测试程序退出太快
	fmt.Printf("counter = %d \n", counter) // counter并不是5000，counter是共享的，并发导致counter会竞争，丢失数据，不是线程安全的
}

// 线程安全的
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { // 类似finally，最后要把锁释放掉
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("counter = %d \n", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup // 比time.Sleep好，我们不能确定协程多久能执行完
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 每个协程加1
		go func() {
			defer func() { // 类似finally，最后要把锁释放掉
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() // 完成一个协程
		}()
	}
	wg.Wait() // 等待协程序结束
	fmt.Printf("counter = %d \n", counter)
}
