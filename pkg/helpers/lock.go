package helpers

import (
	"sync"
)

type MutexWrapper struct {
	mutex sync.Mutex
}

func (mw *MutexWrapper) lock(){
	mw.mutex.Lock()
}

func (mw *MutexWrapper) unLock() {
	mw.mutex.Unlock()
}

// //////////////////////////////

type RWMutexWrapper struct {
	rwMutex sync.RWMutex
}

func (mw *RWMutexWrapper) rLock() {
	mw.rwMutex.RLock()
}

func (mw *RWMutexWrapper) rUnLock() {
	mw.rwMutex.RUnlock()
}

func LockFuncReturnOne[T any](f func() T) {
	m := MutexWrapper{}

	m.lock()
	defer m.unLock()
	f()
}

func LockFuncReturnTwo[T any](f func() (T, T)) {
	m := MutexWrapper{}

	m.lock()
	defer m.unLock()
	f()
}
