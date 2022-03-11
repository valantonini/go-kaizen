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
					time.Sleep(5 * time.Millisecond)
				}
			}
		}()

		time.Sleep(15 * time.Millisecond)
		done <- true
		wg.Wait()

		Is.Equal(count, 2)
	})
}
