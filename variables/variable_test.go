package variables

import (
	"github.com/matryer/is"
	"testing"
)

// Default values for uninitialised variables
func Test_Variables(t *testing.T) {
	Is := is.New(t)

	t.Run("string", func(t *testing.T) {
		var str string
		Is.Equal(str, "")
	})

	t.Run("number", func(t *testing.T) {
		var n int
		Is.Equal(n, 0)
	})

	t.Run("boolean", func(t *testing.T) {
		var b bool
		Is.Equal(b, false)
	})

	t.Run("interface", func(t *testing.T) {
		type interfaceType interface {
			Foo()
		}
		var i interfaceType
		Is.Equal(i, nil)
	})

	t.Run("slice", func(t *testing.T) {
		var slice []int
		Is.Equal(slice, nil)
	})

	t.Run("pointer", func(t *testing.T) {
		var p *int
		Is.Equal(p, nil)
	})

	t.Run("channel", func(t *testing.T) {
		var ch chan int
		Is.Equal(ch, nil)
	})

	t.Run("map", func(t *testing.T) {
		var m map[string]string
		Is.Equal(m, nil)
	})

	t.Run("func", func(t *testing.T) {
		var f func(s string)
		Is.Equal(f, nil)
	})
}
