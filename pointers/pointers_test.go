package pointers

import (
	"github.com/matryer/is"
	"testing"
)

func Test_Pointers(t *testing.T) {
	Is := is.New(t)

	t.Run("pointer addresses", func(t *testing.T) {
		var x = 7
		var p *int      // *int == pointer to int
		p = &x          // &x == address of x
		Is.Equal(*p, 7) // *p == value at address
	})

	t.Run("value at pointer address", func(t *testing.T) {
		x := 1
		p := &x // p, of type *int, points to x
		Is.Equal(*p, 1)

		*p = 2 // equivalent to x = 2
		Is.Equal(x, 2)
	})

	t.Run("pointer parameters are passed by value and do not reassign", func(t *testing.T) {
		target := 777
		originalTargetAddress := &target

		var parameterAddress *int
		reassignPointer := func(p *int) {
			parameterAddress = p
			localNum := 888
			// the address is passed by value to the function. We are reassigning the parameter's address and not the
			// original target's address. If we wanted to reassign the target, we'd need to pass in a **int
			p = &localNum
		}

		reassignPointer(&target)
		Is.True(parameterAddress == originalTargetAddress)
		Is.True(&target == originalTargetAddress)
	})
}
