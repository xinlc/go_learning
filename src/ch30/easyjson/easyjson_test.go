package jsontest

import (
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

// go get -u github.com/mailru/easyjson/...
// ~/go/bin/easyjson -all struct_def.go                // 注意要把项目路径加入到GOPATH中
func TestEasyJson(t *testing.T) {
	// e := new(Employee)
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)

	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}
}
