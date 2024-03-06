package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 自旋锁
type spinLock struct {
	rId int64 `desc:"占有锁的id"`
}

func NewSpinLock() sync.Locker {
	return &spinLock{}
}

func (lock *spinLock) Lock() {
	rId := getRId()
	times := 0
	for !atomic.CompareAndSwapInt64(&lock.rId, nullRid, rId) {
		times++
		fmt.Printf("第%d次自旋", times)
	}
}
func (lock *spinLock) Unlock() {
	rId := getRId()
	if !atomic.CompareAndSwapInt64(&lock.rId, rId, nullRid) {
		panic("You have no enough power to unlock this lock which not belong you!")
	}
}
