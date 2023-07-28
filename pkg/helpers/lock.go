package helpers

import (
	"sync"
)

type MutexWrapper struct {
	mutex sync.Mutex
}

func (mw *MutexWrapper) lock() {
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

func LockFuncOneInOneOut[i any, o any](f func(i) o) func(i) o {
	m := MutexWrapper{}
	return func(iVal i) o {
		m.lock()
		defer m.unLock()
		return f(iVal)
	}
}

func LockFuncOneInTwoOut[i any, o1 any, o2 any](f func(i) (o1, o2)) func(i) (o1, o2) {
	m := MutexWrapper{}
	return func(iVal i) (o1, o2) {
		m.lock()
		defer m.unLock()
		return f(iVal)
	}
}

func LockFuncTwoInTwoOut[i1 any, i2 any, o1 any, o2 any](f func(i1, i2) (o1, o2)) func(i1, i2) (o1, o2) {
	m := MutexWrapper{}

	return func(i1Val i1, i2Val i2) (o1, o2) {
		m.lock()
		defer m.unLock()
		return f(i1Val, i2Val)
	}
}
