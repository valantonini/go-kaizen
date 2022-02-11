package embedding

import (
	"github.com/matryer/is"
	"testing"
)

type OuterStruct struct {
	EmbeddedStruct
}

func (OuterStruct) OuterStructReceiver() string {
	return "foo"
}

type EmbeddedStruct struct {
	EmbeddedStructField string
}

func (EmbeddedStruct) EmbeddedStructReceiver() string {
	return "bar"
}

func Test_Embedding(t *testing.T) {
	Is := is.New(t)

	t.Run("embedded structs promote receivers to parent", func(t *testing.T) {
		// All public receivers on EmbeddedStruct will be promoted to OuterStruct. Care must be taken as this can result in exposing
		// fields that should remain internal (such as mutexes)
		foo := new(OuterStruct)

		Is.Equal(foo.EmbeddedStructReceiver(), "bar")
	})

	t.Run("embedded structs receivers can be reached via the embedded type", func(t *testing.T) {
		foo := new(OuterStruct)

		Is.Equal(foo.EmbeddedStruct.EmbeddedStructReceiver(), "bar")
	})

	t.Run("embedded structs promote fields", func(t *testing.T) {
		foo := OuterStruct{EmbeddedStruct{"embedded_field_value"}}

		Is.Equal(foo.EmbeddedStructField, "embedded_field_value")
	})

	t.Run("embedded structs satisfy interfaces", func(t *testing.T) {
		type Embedder interface {
			EmbeddedStructReceiver() string
		}

		var e Embedder
		e = OuterStruct{}

		Is.Equal(e.EmbeddedStructReceiver(), "bar")
	})
}
