package variables

import (
	"github.com/matryer/is"
	"testing"
)

// Default values for uninitialised variables
func Test_Variables(t *testing.T) {
	is := is.New(t)

	t.Run("string", func(t *testing.T) {
		var str string
		is.Equal(str, "")
	})

	t.Run("number", func(t *testing.T) {
		var n int
		is.Equal(n, 0)
	})

	t.Run("boolean", func(t *testing.T) {
		var b bool
		is.Equal(b, false)
	})

	t.Run("interface", func(t *testing.T) {
		type interfaceType interface {
			Foo()
		}
		var i interfaceType
		is.Equal(i, nil)
	})

	t.Run("slice", func(t *testing.T) {
		var slice []int
		is.Equal(slice, nil)
	})

	t.Run("pointer", func(t *testing.T) {
		var p *int
		is.Equal(p, nil)
	})

	t.Run("channel", func(t *testing.T) {
		var ch chan int
		is.Equal(ch, nil)
	})

	t.Run("map", func(t *testing.T) {
		var m map[string]string
		is.Equal(m, nil)
	})

	t.Run("func", func(t *testing.T) {
		var f func(s string)
		is.Equal(f, nil)
	})
}
