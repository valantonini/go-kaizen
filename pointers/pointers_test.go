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
}
