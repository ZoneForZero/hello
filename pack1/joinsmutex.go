package pack1

import (
	"sync"
	"sync/atomic"
)

// 可重入锁
type locker struct {
	lockNum int64      `desc:"锁次数"`
	lockRId int64      `desc:"持有的id"`
	cond    *sync.Cond `desc:"释放条件"` // implements goroutine blocking and waking with sync.Cond
}

func New() sync.Locker {
	return &locker{}
}

// 自我线程再次锁定
func (l *locker) tryOwnerAcquire(gid int64) bool {
	// 乐观锁的实现
	if atomic.CompareAndSwapInt64(&l.lockRId, gid, gid) {
		atomic.AddInt64(&l.lockNum, 1)
		return true
	}
	return false
}

// 占用无人持有的锁
func (l *locker) tryNoneAcquire(gid int64) bool {
	// CAS none -> gid, 若成功则代表原本没有线程持有锁，当前线程获取到了锁
	if atomic.CompareAndSwapInt64(&l.lockRId, nullRid, gid) {
		// 这里跟上一步就不是原子操作了
		atomic.AddInt64(&l.lockNum, 1)
		return true
	}
	return false
}

// 自旋等待
func (l *locker) acquireSlow(gid int64) {
	l.cond.L.Lock()
	// 若当前还是有其他线程持有锁
	for atomic.LoadInt64(&l.lockRId) != nullRid {
		// 继续等待唤醒
		l.cond.Wait()
	}
	{ // 若当前没有线程持有锁, 则持有锁
		atomic.SwapInt64(&l.lockRId, gid)
		atomic.SwapInt64(&l.lockNum, 1)
	}
	l.cond.L.Unlock()
}

// Lock locks the locker with reentrant mode
func (l *locker) Lock() {
	gid := getRId()
	if l.tryOwnerAcquire(gid) {
		return
	}
	if l.tryNoneAcquire(gid) {
		return
	}
	l.acquireSlow(gid)
}

// Unlock unlocks the locker using reentrant mode
func (l *locker) Unlock() {
	lockRId := atomic.LoadInt64(&l.lockRId)
	if lockRId == nullRid {
		panic("unlock of unlocked reentrantLocker")
	}
	// 判断是自己，就可以解锁了
	if lockRId == getRId() {
		if atomic.LoadInt64(&l.lockNum) == 0 {
			panic("unlocks more than locks")
		}
		l.lockNum--
		if l.lockNum == 0 {
			l.lockRId = nullRid
			// signal one and only one goroutine
			l.cond.Signal()
		}
		return
	}
	panic("unlock by different goroutine")
}
