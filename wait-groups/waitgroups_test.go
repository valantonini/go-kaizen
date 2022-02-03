package wait_groups

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// https://pkg.go.dev/sync#WaitGroup
// A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
// A WaitGroup must not be copied after first use.

// DoRequest returns the value as string after a delay relative to it's value
func DoRequest(delay int) string {
	time.Sleep(time.Duration(delay*5) * time.Millisecond)
	return strconv.Itoa(delay)
}

func ExampleDoRequest() {
	var wg sync.WaitGroup

	for i := 3; i > 0; i-- {
		wg.Add(1)

		go func(delay int) {
			// decrement wait group count on scope exit
			defer wg.Done()
			response := DoRequest(delay)
			fmt.Println(response)
		}(i) // don't use loop variables in go routine as it can have unexpected side effects
	}

	// block unit all go routines finish
	wg.Wait()

	// Output: 1
	// 2
	// 3
}
