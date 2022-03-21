package recover

import (
	"fmt"
	"github.com/matryer/is"
	"sync"
	"testing"
)

func panicOnEven(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	if n%2 == 0 {
		panic("even number")
	}
}

// Test_RecoverGoRoutine shows how to guard against panic in go routines by
// delegating dangerous work to a function which is guarded by a defer/recover
func Test_RecoverGoRoutine(t *testing.T) {
	Is := is.New(t)

	success := 0

	var wg sync.WaitGroup
	wg.Add(1)

	go func(iterationsCompleted *int) {
		for i := 1; i <= 3; i++ {
			panicOnEven(i)
			*iterationsCompleted++
		}
		wg.Done()
	}(&success)

	wg.Wait()

	Is.Equal(success, 3)
}
