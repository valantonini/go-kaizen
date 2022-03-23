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

	t.Run("slices can be declared with a length and a capacity of the backing array", func(t *testing.T) {
		source := make([]int, 2, 3)
		source[0] = 0
		source[1] = 1

		// will panic, length is only 2 even though the backing array is 3
		source = append(source, 2)

		Is.Equal(len(source), 3)
		Is.Equal(cap(source), 3)

		subslice := source[0:3]
		Is.Equal(subslice[0], 0)
		Is.Equal(subslice[1], 1)
		Is.Equal(subslice[2], 2)

		// change first value of original slice
		source[0] = 4
		Is.Equal(subslice[0], 4)

		// this will cause backing array to be reallocated (doubles size when capacity is reached)
		subslice = append(subslice, 6)
		Is.Equal(cap(subslice), 6)

		// source and subslice now use 2 different backing arrays. this will only update the source
		source[0] = 7
		Is.Equal(subslice[0], 4)
	})

	t.Run("a nil slice can be appended to", func(t *testing.T) {
		var items []string

		Is.Equal(items, nil)

		items = append(items, "foo", "bar")

		Is.Equal(items[0], "foo")
		Is.Equal(items[1], "bar")
	})

	t.Run("a nil slice will panic if indexed", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				Is.Equal(r.(error).Error(), "runtime error: index out of range [0] with length 0")
			}
		}()

		var items []string
		items[0] = "foo"

		// will jump to recover instead of reaching here
		Is.Fail()
	})
}
