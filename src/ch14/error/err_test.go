package error_test

import (
	"errors"
	"fmt"
	"testing"
)

// Go没有异常机制
// 通过多返回值判断
// 及早失败，避免嵌套，错误判断写在前面
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("22222222")
		}
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
