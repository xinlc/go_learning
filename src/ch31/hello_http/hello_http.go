package main

import (
	"fmt"
	"net/http"
	"time"
)

// go run hello_http.go
// 路由规则:
// URL 分两种，末尾是 /: 表示一个子树，后面可以跟其他子路径；末尾不是/，表示一个叶子，固定路径。
// 采用最长匹配原则，如果有多个匹配，会采用匹配路径最长的那个进行处理
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
