package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象
}

// 创建对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj) // 池大小
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{} // 这里放个空结构，可以用别的，连接之类的。。。
	}
	return &objPool
}

// 获取对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) { // 想这个池可以放不同的对象，可以用obj interface{}, 但不建议这样做，什么池就放什么对象
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

// 放回
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
