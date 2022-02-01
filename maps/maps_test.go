package maps

import (
	"github.com/matryer/is"
	"testing"
)

func Test_Maps(t *testing.T) {
	Is := is.New(t)

	t.Run("missing key for value type will return default value on read", func(t *testing.T) {
		strLookup := map[int]string{
			1: "Foo",
			2: "Bar",
		}

		Is.Equal(strLookup[999], "")
	})

	t.Run("missing key for reference types will return nil on read", func(t *testing.T) {
		type vector2 struct {
			x int
			y int
		}

		vectorLookup := map[int]*vector2{
			1: {1, 2},
			2: {3, 4},
		}

		Is.Equal(vectorLookup[999], nil)
	})

	t.Run("a nil map wont error on reading", func(t *testing.T) {
		var strLookup map[int]string

		Is.Equal(strLookup[999], "")
	})

	t.Run("a nil map will explode on writing", func(t *testing.T) {
		var strLookup map[int]string

		defer func() {
			if err := recover(); err != nil {
				Is.Equal(strLookup[1], "")
			}
		}()

		strLookup[1] = "foo"

		// this should never be reached as the assignment above will cause panic that calls the deferred function and
		// returns
		Is.Fail()
	})
}
