package channels

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

func Test_Channels(t *testing.T) {
	Is := is.New(t)

	t.Run("select + returning channels can be used to perform concurrent requests and continue on first response", func(t *testing.T) {
		doRequest := func(delay int, response string) chan string {
			c := make(chan string)
			go func() {
				time.Sleep(time.Duration(delay) * time.Millisecond)
				c <- response
			}()
			return c
		}

		got := ""
		select {
		case slowRequest := <-doRequest(5, "slow"):
			got = slowRequest
		case fastRequest := <-doRequest(2, "fast"):
			got = fastRequest
		}

		Is.Equal(got, "fast")
	})

	t.Run("ranging over channels", func(t *testing.T) {
		queue := make(chan int, 2)
		queue <- 1
		queue <- 1

		// it’s possible to close a non-empty channel but still have the remaining values be received.
		close(queue)

		count := 0

		// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
		for elem := range queue {
			count += elem
		}

		Is.Equal(count, 2)
	})
}
