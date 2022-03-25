package patterns

import (
	"context"
	"golang.org/x/sync/errgroup"

	"github.com/matryer/is"
	"testing"
	"time"
)

func DoAsyncTask() <-chan string {
	result := make(chan string, 1)

	go func() {
		data := DoSlowTask()
		result <- data
	}()

	return result
}

func DoSlowTask() string {
	time.Sleep(15 * time.Millisecond)
	return "completed"
}

func Test_Futures(t *testing.T) {
	Is := is.New(t)
	t.Run("a future returns a channel immediately that is populated from a different go routine", func(t *testing.T) {

		c := DoAsyncTask()

		result := <-c

		Is.Equal(result, "completed")
	})

	t.Run("wait for channels in method signature to prevent blocking", func(t *testing.T) {

		doWork := func(p1, p2 string) {
			Is.Equal(p1, "completed")
			Is.Equal(p2, "completed")
		}

		c1 := DoAsyncTask()
		c2 := DoAsyncTask()

		doWork(<-c1, <-c2)
	})

	t.Run("futures with error groups", func(t *testing.T) {
		var result1, result2 string
		g, ctx := errgroup.WithContext(context.TODO())

		g.Go(func() error {
			result1 = DoSlowTask()
			return nil
		})

		g.Go(func() error {
			result2 = DoSlowTask()
			return nil
		})

		Is.True(ctx != nil)

		err := g.Wait()

		Is.NoErr(err)
		Is.Equal(result1, "completed")
		Is.Equal(result2, "completed")
	})
}
