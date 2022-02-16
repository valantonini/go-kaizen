package reflection

import (
	"github.com/matryer/is"
	"reflect"
	"testing"
)

func Test_Reflection(t *testing.T) {
	Is := is.New(t)

	t.Run("retrieving the number of fields will scoop up internal fields", func(t *testing.T) {
		person := struct {
			Name       string
			Occupation string
			age        int
		}{
			"arnold",
			"engineer",
			26,
		}

		r := reflect.ValueOf(person)

		Is.Equal(r.NumField(), 3)
	})
	t.Run("addressing fields by order", func(t *testing.T) {
		person := struct {
			Name string
		}{
			"arnold",
		}

		r := reflect.ValueOf(person)

		// r.Field() returns reflect.Value(arnold)
		got := r.Field(0).String() // DANGER: will panic if no fields or not string
		Is.Equal(got, "arnold")
	})

	t.Run("asserting type", func(t *testing.T) {
		person := struct {
			Name string
		}{
			"arnold",
		}

		r := reflect.ValueOf(person)

		Is.True(r.Field(0).Kind() == reflect.String)
	})

	t.Run("using reflection to dereference a pointer", func(t *testing.T) {
		type occupation struct {
			title string
		}
		person := struct {
			Name       string
			Occupation *occupation
		}{
			"arnold",
			&occupation{"engineer"},
		}

		r := reflect.ValueOf(person)

		Is.True(r.Field(1).Kind() == reflect.Ptr)
		Is.Equal(r.Field(1).Elem().String(), "engineer")
	})
}
