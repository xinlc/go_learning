package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者 消费者
// channel 可以关闭，实现多个消费者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 关闭 channel
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// for i := 0; i < 10; i++ {
		// 	if data, ok := <-ch; ok { // ok==false 表示channel关闭
		// 		fmt.Println(data)
		// 	} else {
		// 		break
		// 	}
		// }
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
