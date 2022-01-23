package equality

import (
	"fmt"
	"github.com/matryer/is"
	"testing"
)

func Test_Equality(t *testing.T) {
	Is := is.New(t)

	t.Run("when nil isn't nil", func(t *testing.T) {
		// https://www.calhoun.io/when-nil-isnt-equal-to-nil/

		isNil := func(target interface{}) bool {
			return t == nil
		}

		var arr map[string]string = nil
		isNil(arr)
		Is.Equal(arr, nil)

		// in Go has two basic pieces of information; the type of the pointer, and the value it points to
		mapTypeInfo := fmt.Sprintf("(%T, %v)\n", arr, arr)
		Is.Equal(mapTypeInfo, "([]string, [])")

		nilTypeInfo := fmt.Sprintf("(%T, %v)\n", nil, nil)
		Is.Equal(nilTypeInfo, "(<nil>, <nil>)")

		// isNil is incorrectly implemented and will incorrectly report false
		Is.Equal(isNil(arr), false)

		// What we are really comparing is both the values AND the types. That is, we are not just comparing the value
		// stored in a with the nil value; we are also comparing their types.

		// ([]string, []) != (<nil>, <nil>)
	})
}
