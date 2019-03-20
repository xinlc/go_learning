package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// go 不支持继承，而是用复合
// 复合S
// type Dog struct {
// 	p *Pet
// }

// func (d *Dog) Speak() {
// 	fmt.Println("Wang!")
// 	// d.p.Speak()
// }

// func (d *Dog) SpeakTo(host string) {
// 	// d.p.SpeakTo(host)
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

// func TestDog(t *testing.T) {
// 	dog := new(Dog)
// 	dog.SpeakTo("Chao")
// }
// 复合E

type Dog struct {
	Pet
}

func TestDog(t *testing.T) {
	var dog = new(Dog)
	dog.SpeakTo("Chao1")
}
