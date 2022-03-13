package equality

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func Test_Equality(t *testing.T) {
	Is := is.New(t)

	t.Run("test slice equality using reflect.DeepEqual", func(t *testing.T) {
		slice1 := []int{3, 4, 5}
		slice2 := []int{3, 4, 5}

		// https://pkg.go.dev/reflect#DeepEqual
		Is.True(reflect.DeepEqual(slice1, slice2) == true)
	})

	t.Run("test struct value equality is based on values not reference", func(t *testing.T) {
		type vector2 struct {
			x int
			y int
		}

		v1 := vector2{3, 14}
		v2 := vector2{3, 14}

		Is.True(v1 == v2)
	})

	t.Run("test struct value equality will use pointer address and not value at pointer unless reflect.DeepEqual is used", func(t *testing.T) {
		type node struct {
			value  string
			parent *node
		}

		n1 := node{"foo", &node{"bar", nil}}
		n2 := node{"foo", &node{"bar", nil}}

		// not equal as address of parent nodes are different unless reflect.DeepEqual is used
		Is.Equal(n1 == n2, false)
		Is.True(reflect.DeepEqual(n1, n2))

		parent := node{"bar", nil}
		n1 = node{"foo", &parent}
		n2 = node{"foo", &parent}

		Is.True(n1 == n2)
		Is.True(reflect.DeepEqual(n1, n2))
	})

	t.Run("uses address and not value for structural equality", func(t *testing.T) {
		type wrapper struct {
			value *string
		}

		//same value, different address
		val1 := "foo"
		val2 := "foo"

		n1 := wrapper{&val1}
		n2 := wrapper{&val2}

		Is.True(n1 != n2)

		// same value, same address
		n2 = wrapper{&val1}
		Is.True(n1 == n2)
	})
}
