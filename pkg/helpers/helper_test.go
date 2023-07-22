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
