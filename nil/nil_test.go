package nil

import (
	"fmt"
	"github.com/matryer/is"
	"testing"
)

func Test_Nil(t *testing.T) {
	Is := is.New(t)

	t.Run("when nil isn't nil", func(t *testing.T) {
		// https://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/
		// https://www.calhoun.io/when-nil-isnt-equal-to-nil/

		isNil := func(target interface{}) bool {
			return t == nil
		}

		var arr []string = nil
		Is.Equal(arr, nil)

		// in Go has two basic pieces of information; the type of the pointer, and the value it points to
		arrTypeInfo := fmt.Sprintf("(%T, %v)", arr, arr)
		fmt.Print(arrTypeInfo)
		Is.Equal(arrTypeInfo, "([]string, [])")

		nilTypeInfo := fmt.Sprintf("(%T, %v)", nil, nil)
		Is.Equal(nilTypeInfo, "(<nil>, <nil>)")

		// isNil is incorrectly implemented and will incorrectly report false
		Is.Equal(isNil(arr), false)

		// What we are really comparing is both the values AND the types. That is, we are not just comparing the value
		// stored in a with the nil value; we are also comparing their types.

		// ([]string, []) != (<nil>, <nil>)
	})
}
