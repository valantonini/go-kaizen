package recover

import (
	"github.com/matryer/is"
	"testing"
)

// https://gobyexample.com/recover
// Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution instead.

func Test_Recover(t *testing.T) {
	Is := is.New(t)

	// recover must be called within a deferred function. When the enclosing function panics, the defer will activate
	// and a recover call within it will catch the panic.
	defer func() {
		if r := recover(); r != nil {
			// The return value of recover is the error raised in the call to panic.
			Is.Equal(r, "error has occurred")
		}
	}()

	panic("error has occurred")
}
