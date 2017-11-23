package semaphore

import (
	"testing"
	"fmt"
)

func BenchmarkSemaphore2_Acquire(b *testing.B) {
	sem := New2(b.N)

	for i := 0; i < b.N; i++ {
		sem.Acquire(nil)
	}

	if sem.GetCount() != sem.GetLimit() {
		b.Error(fmt.Printf("semaphore must have count = %v, but has %v", sem.GetLimit(), sem.GetCount()))
	}
}

func BenchmarkSemaphore2_Acquire_Release(b *testing.B) {
	sem := New2(b.N)

	for i := 0; i < b.N; i++ {
		sem.Acquire(nil)
		sem.Release()
	}

	if sem.GetCount() != 0 {
		b.Error("semaphore must have count = 0")
	}
}
