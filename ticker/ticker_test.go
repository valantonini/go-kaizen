package ticker

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

// Ticker is useful for performing tasks at regular intervals / polling

func Test_Ticker(t *testing.T) {
	Is := is.New(t)
	done := make(chan bool) // to cancel the go routine
	count := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for {
			select {
			case <-done:
				return // exit the go routine
			case <-ticker.C:
				count++
			}
		}
	}()

	time.Sleep(35 * time.Millisecond)
	done <- true
	Is.Equal(count, 3)
}
