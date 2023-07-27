package helpers

import (
	"testing"
	"reflect"
	"time"
)

func TestRandomString(t *testing.T) {
	got := RandomString(10)
	if len(got) != 10 {
		t.Errorf("RandomString(%d) = %s; want 10 characters", len(got), got)
	}

	if reflect.TypeOf(got).Kind()	!= reflect.String {
		t.Errorf("RandomString(%d) = %s; want only letters", len(got), got)
	}
}

func TestParallelize(t *testing.T) {
	done := make(chan struct{})

	func1 := func() {
		time.Sleep(1 * time.Second)
	}

	func2 := func() {
		time.Sleep(1 * time.Second)
	}

	go func() {
		Parallelize(func1, func2)
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(2 *time.Second):
		t.Error("Parallelize(): Timeout waiting to complete, maximum 2 seconds")
	}
}

// Lock

func TestLockFuncOneInOneOut(t *testing.T) {
	var got int

	func1 := func(i int) int {
		return i
	}

	lockedFunc := LockFuncOneInOneOut(func1)

	go func() {
		got = lockedFunc(1)
	}()

	time.Sleep(1 * time.Second)

	if got != 1 {
		t.Errorf("LockFuncOneInOneOut(): got %d; want 1", got)
	}
}

func TestLockFuncTwoInTwoOut(t *testing.T) {
	var got1, got2 int

	func1 := func(i1 int, i2 int) (int, int) {
		return i1, i2
	}

	lockedFunc := LockFuncTwoInTwoOut(func1)

	go func() {
		got1, got2 = lockedFunc(1, 2)
	}()

	time.Sleep(1 * time.Second)

	if got1 != 1 || got2 != 2 {
		t.Errorf("LockFuncTwoInTwoOut(): got (%d, %d); want (1, 2)", got1, got2)
	}
}
