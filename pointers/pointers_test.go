package pointers

import (
	"github.com/valantonini/go-kaizen/assert"
	"testing"
)

func Test_Pointers(t *testing.T) {
	t.Run("pointer addresses", func(t *testing.T) {
		var x int
		var p *int          // *int == pointer to int
		p = &x              // &x == address of x
		assert.NotNil(t, p) // *p == value at address
	})

	t.Run("value at pointer address", func(t *testing.T) {
		x := 1
		p := &x // p, of type *int, points to x
		assert.Equal(t, *p, 1)

		*p = 2 // equivalent to x = 2
		assert.Equal(t, x, 2)
	})
}
