package numbers

import (
	"github.com/matryer/is"
	"testing"
)

func Test_Numbers(t *testing.T) {
	Is := is.New(t)

	t.Run("numbers can include _ separators for clarity", func(t *testing.T) {
		number := 10_000
		Is.Equal(number, 10000)
	})

	t.Run("numbers preceded with 0 are octal", func(t *testing.T) {
		number := 010

		Is.Equal(number, 8)

		// useful for *nix file permissions e.g.
		// os.OpenFile("foo", os.O_RDONLY, 0644)
	})

	t.Run("0o can be used to make octal clearer", func(t *testing.T) {
		number := 0o11

		Is.Equal(number, 9)
	})

	t.Run("0b can be used for binary", func(t *testing.T) {
		number := 0b11

		Is.Equal(number, 3)
	})

	t.Run("0x can be used for binary", func(t *testing.T) {
		number := 0x11

		Is.Equal(number, 17)
	})

	t.Run("i can be used for imaginary number", func(t *testing.T) {
		// https://en.wikipedia.org/wiki/Imaginary_number
		number := 5i

		got := number * number

		Is.Equal(real(got), float64(-25))
	})
}
