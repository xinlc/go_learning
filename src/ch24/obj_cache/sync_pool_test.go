package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Pool 不是同步池，可以理解成sync.cache
// 适用于通过复用，降低复杂对象的创建和GC代价
// 携程安全，会有锁的开销
// 生命周期受GC影响，不适合于做连接池等，需自己管理生命周期的资源的池花

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	// runtime.GC() // GC 会清楚sync.pool中缓存的对象, 一般不用人为触发
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	// 先输出3个100， 没有了在走创建
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
