package context

import (
	"context"
	"github.com/matryer/is"
	"sync"
	"testing"
	"time"
)

func sleepAndSet(context context.Context, wg *sync.WaitGroup, ms int, value *string) {
	select {
	case <-time.After(time.Duration(ms) * time.Millisecond):
		*value = "completed"
	case <-context.Done():
		*value = "cancelled from context"
	}

	wg.Done()
}

// https://go.dev/blog/context

func Test_Context(t *testing.T) {
	Is := is.New(t)

	t.Run("should finish execution before context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel() // need to call cancel to release resources

		var wg sync.WaitGroup
		result := ""

		wg.Add(1)
		go sleepAndSet(ctx, &wg, 10, &result)

		wg.Wait()
		Is.Equal(result, "completed")

	})

	t.Run("should cancel execution via context timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		defer cancel() // need to call cancel to release resources

		var wg sync.WaitGroup
		result := ""

		wg.Add(1)
		go sleepAndSet(ctx, &wg, 10, &result)

		wg.Wait()
		Is.Equal(result, "cancelled from context")
	})

	t.Run("withcancel can be used to derive a context that can be cancelled sooner", func(t *testing.T) {

		childValue := ""
		parentValue := ""

		// child will time out after 10ms
		child := func(ctx context.Context) {
			select {
			case <-ctx.Done():
				childValue = "child cancelled"
			case <-time.After(10 * time.Millisecond):
				childValue = "child completed"
			}

		}

		// parent will invoke child and cancel after 5ms
		parent := func(ctx context.Context) {
			callingCtx, cancel := context.WithCancel(ctx)
			time.AfterFunc(5*time.Millisecond, cancel)
			go child(callingCtx)

			select {
			case <-ctx.Done():
				parentValue = "parent cancelled"
			case <-time.After(10 * time.Millisecond):
				parentValue = "parent completed"
			}
		}

		c, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
		defer cancel()

		parent(c)

		Is.Equal(childValue, "child cancelled")
		Is.Equal(parentValue, "parent completed")
	})

	t.Run("withcancel can be used to derive a context that can cancels in step with parent", func(t *testing.T) {

		childValue := ""
		parentValue := ""

		// child will time out after 10ms
		child := func(ctx context.Context) {
			select {
			case <-ctx.Done():
				childValue = "child cancelled"
			case <-time.After(10 * time.Millisecond):
				childValue = "child completed"
			}

		}

		// parent will invoke child and cancel after 5ms
		parent := func(ctx context.Context) {
			callingCtx, cancel := context.WithCancel(ctx)
			time.AfterFunc(5*time.Millisecond, cancel)
			go child(callingCtx)

			select {
			case <-ctx.Done():
				parentValue = "parent cancelled"
			case <-time.After(10 * time.Millisecond):
				parentValue = "parent completed"
			}
		}

		c, cancel := context.WithTimeout(context.Background(), 4*time.Millisecond)
		defer cancel()

		// invoke the parent and cancel, it will cancel after 4ms cancelling it's child derived contexts too
		parent(c)

		Is.Equal(childValue, "child cancelled")
		Is.Equal(parentValue, "parent cancelled")
	})

	t.Run("withcancel can be used to derive a context", func(t *testing.T) {

		childValue := ""
		parentValue := ""

		// child will time out after 4ms
		child := func(ctx context.Context) {
			select {
			case <-ctx.Done():
				childValue = "child cancelled"
			case <-time.After(4 * time.Millisecond):
				childValue = "child completed"
			}

		}

		// parent will invoke child and cancel after 8ms
		parent := func(ctx context.Context) {
			callingCtx, cancel := context.WithCancel(ctx)
			time.AfterFunc(8*time.Millisecond, cancel)
			go child(callingCtx)

			select {
			case <-ctx.Done():
				parentValue = "parent cancelled"
			case <-time.After(10 * time.Millisecond):
				parentValue = "parent completed"
			}
		}

		c, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
		defer cancel()

		parent(c)

		Is.Equal(childValue, "child completed")
		Is.Equal(parentValue, "parent completed")
	})
}
