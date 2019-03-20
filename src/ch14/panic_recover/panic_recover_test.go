package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

// panic vs os.Exit
// 1. os.Exit 退出不会调用derfer指定的函数
// 2. os.Exit 退出时不会输出当前调用栈信息

// recover 可以恢复错误，类似java try{} catch(Throwable)
// 但revocer很危险，可能导致僵尸服务进程，导致health check 失效。
// Let it Crash! 往往是我们恢复不确定性错误的最好方法。
func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
	// fmt.Println("End")

}
