package strings

import (
	"github.com/matryer/is"
	"strconv"
	"strings"
	"testing"
)

func Test_Strings(t *testing.T) {
	Is := is.New(t)

	t.Run("converting integers to string", func(t *testing.T) {
		number := 123

		// itoa == integer to ascii
		str := strconv.Itoa(number)

		Is.Equal(str, "123")
	})

	t.Run("parsing strings to integers", func(t *testing.T) {
		str := "123"

		// atoi == ascii to integer
		number, err := strconv.Atoi(str)

		Is.NoErr(err)
		Is.Equal(number, 123)
	})

	t.Run("can convert to and from []byte", func(t *testing.T) {
		foo := "foo"
		bytes := []byte(foo)
		Is.Equal(bytes, []byte{102, 111, 111})

		str := string(bytes)
		Is.Equal(str, "foo")
	})

	t.Run("stringbuilder for efficiently concatenating strings", func(t *testing.T) {
		var sb strings.Builder

		sb.WriteString("foo")
		sb.WriteString("bar")

		Is.Equal(sb.String(), "foobar")
	})
}
