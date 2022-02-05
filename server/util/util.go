package util

import (
	"sync/atomic"
)

type Lock struct {
	val LockVal
}

type LockVal = int32

const (
	locked   LockVal = 1
	unlocked LockVal = 0
)

func NewLock() *Lock {
	return &Lock{
		val: unlocked,
	}
}

func (l *Lock) TryLock() (succ bool) {
	return atomic.CompareAndSwapInt32(&l.val, unlocked, locked)
}

func (l *Lock) Unlock() {
	atomic.SwapInt32(&l.val, unlocked)
}
