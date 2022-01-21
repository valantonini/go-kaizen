package pointers

import (
	"github.com/matryer/is"
	"github.com/valantonini/go-kaizen/assert"
	"testing"
)

func Test_Pointers(t *testing.T) {
	is := is.New(t)
	t.Run("pointer addresses", func(t *testing.T) {
		var x int = 7
		var p *int      // *int == pointer to int
		p = &x          // &x == address of x
		is.Equal(*p, 7) // *p == value at address
	})

	t.Run("value at pointer address", func(t *testing.T) {
		x := 1
		p := &x // p, of type *int, points to x
		assert.Equal(t, *p, 1)

		*p = 2 // equivalent to x = 2
		assert.Equal(t, x, 2)
	})
}
