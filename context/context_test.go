package context

import (
	"context"
	"github.com/matryer/is"
	"sync"
	"testing"
	"time"
)

func sleepAndSet(context context.Context, wg *sync.WaitGroup, value *string) {
	select {
	case <-time.After(10 * time.Millisecond):
		*value = "completed"
	case <-context.Done():
		*value = "cancelled from context"
	}

	wg.Done()
}

func Test_Context(t *testing.T) {
	Is := is.New(t)

	t.Run("should finish execution before context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel() // need to call cancel to release resources

		var wg sync.WaitGroup
		result := ""

		wg.Add(1)
		go sleepAndSet(ctx, &wg, &result)

		wg.Wait()
		Is.Equal(result, "completed")

	})

	t.Run("should cancel execution via context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		defer cancel() // need to call cancel to release resources

		var wg sync.WaitGroup
		result := ""

		wg.Add(1)
		go sleepAndSet(ctx, &wg, &result)

		wg.Wait()
		Is.Equal(result, "cancelled from context")
	})
}
