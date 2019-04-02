package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 单例模式
type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletionObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Ob")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletionObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
