package mutexes

import (
	"sync"
	"testing"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func Test_Mutexes(t *testing.T) {
	t.Run("run safely", func(t *testing.T) {
		concurrency := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(concurrency)

		for i := 0; i < concurrency; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		Assert(t, &counter, concurrency)
	})
}

// this code is unsafe:
// func Assert(t testing.T, got Counter, want int) {
// }
// go ver ./mutexes/mutexes_test.go
// mutexes/mutexes_test.go:42:15: Assert passes lock by value: testing.T contains testing.common contains sync.RWMutex
// mutexes/mutexes_test.go:42:30: Assert passes lock by value: command-line-arguments.Counter contains sync.Mutex
// A Mutex must not be copied after first use.
// https://pkg.go.dev/sync#Mutex
// A Mutex must not be copied after first use.
// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
// When we pass our Counter (by value) to assertCounter it will try and create a copy of the mutex.

func Assert(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.value != want {
		t.Errorf("got %d want %d", got.value, want)
	}
}
