package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 通常为力避免内存拷贝我们使用这种定义方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	fmt.Println()
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 在实例对应访问被调用时，实例成员会进行复制
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针，相当于 e:= &Employee{}, 但不用->符号调用，
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Printf("e is %T", e)
	fmt.Println()
	fmt.Printf("e2 is %T", e2)
	fmt.Println()
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	fmt.Println()
	fmt.Println(e.String())
}
