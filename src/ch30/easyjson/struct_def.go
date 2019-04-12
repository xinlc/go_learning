package jsontest

type BasicInfo struct {
	Name string
	Age  int
}

type JobInfo struct {
	Skills []string
}

type Employee struct {
	BasicInfo BasicInfo
	JobInfo   JobInfo
}
