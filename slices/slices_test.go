package slices

import (
	"testing"

	"github.com/matryer/is"
)

func Test_Slices(t *testing.T) {
	Is := is.New(t)
	t.Run("slicing via indices uses same backing array", func(t *testing.T) {
		source := make([]int, 3)
		source[0] = 0
		source[1] = 1

		subslice := source[1:3]
		Is.Equal(subslice[0], 1)
		Is.Equal(subslice[1], 0) //default uninitialized value

		source[2] = 2
		Is.Equal(subslice[0], 1)
		Is.Equal(subslice[1], 2) //value from source[2] assignment above
	})
}
