package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil { // 尝试放置超出池大小
	// 	fmt.Println(err)
	// }

	// for i := 0; i < 11; i++ {
	for i := 0; i < 10; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T\n", v)
			// if err := pool.ReleaseObj(v); err != nil {
			// 	fmt.Println(err)
			// }
		}
	}
	fmt.Println("Done")
}
