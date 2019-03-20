package client

import (
	series "ch15/series"
	"fmt"
	"testing"
)

// package
// 1. 基本复用模块单元
//     以首字母大写来表明可被包外代码访问
// 2. 代码的package可以和所在的目录不一致
// 3. 同一目录的Go代码的package要保持一致
// 4. 要在PATH中加入模块路径，e.g. /Users/leo/go:/Users/leo/Documents/github/go_learning

// 通过go get来获取远程依赖
// go get -u 强制从网络更新远程依赖
// 注意代码在github上的组织形式，以适应 go get 直接以代码路径开始，不要用src
// 可以用依赖管理工具，godep glide 等，利用go vendor目录做依赖
func TestPackage(t *testing.T) {
	fmt.Println(series.GetFibonacciSerie(10))
}
