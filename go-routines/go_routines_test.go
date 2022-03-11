package go_routines

import (
	"sync"
	"testing"
	"time"

	"github.com/matryer/is"
)

func Test_GoRoutines(t *testing.T) {
	Is := is.New(t)

	t.Run("closing a go routine", func(t *testing.T) {
		done := make(chan bool)
		count := 0

		// wait group to prove channel has closed
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			for {
				select {
				case _ = <-done:
					wg.Done()
					return
				default:
					count++
					time.Sleep(1 * time.Millisecond)
				}
			}
		}()

		time.Sleep(5 * time.Millisecond)
		done <- true
		wg.Wait()

		current := count
		Is.True(count > 0)

		// ensure the counter isn't incrementing
		time.Sleep(5 * time.Millisecond)
		Is.Equal(count, current)
	})
}
