package variables

import (
	"fmt"
	"github.com/valantonini/go-kaizen/assert"
	"testing"
)

// Default values for uninitialised variables
func Test_Variables(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var str string
		assert.Equal(t, str, "")
	})

	t.Run("number", func(t *testing.T) {
		var n int
		assert.Equal(t, n, 0)
	})

	t.Run("boolean", func(t *testing.T) {
		var b bool
		fmt.Println(b == false)
	})

	t.Run("interface", func(t *testing.T) {
		type interfaceType interface {
			Foo()
		}
		var i interfaceType
		assert.Nil(t, i)
	})

	t.Run("slice", func(t *testing.T) {
		var slice []int
		assert.Nil(t, slice)
	})

	t.Run("pointer", func(t *testing.T) {
		var p *int
		assert.Nil(t, p)
	})

	t.Run("channel", func(t *testing.T) {
		var ch chan int
		assert.Nil(t, ch)
	})

	t.Run("map", func(t *testing.T) {
		var m map[string]string
		assert.Nil(t, m)
	})

	t.Run("func", func(t *testing.T) {
		var f func(s string)
		assert.Nil(t, f)
	})
}
