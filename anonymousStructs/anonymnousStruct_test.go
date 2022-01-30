package anonymousStructs

import (
	"bytes"
	"encoding/json"
	"github.com/matryer/is"
	"testing"
)

func Test_AnonymousStructs(t *testing.T) {
	Is := is.New(t)

	t.Run("it serializes to JSON correctly", func(t *testing.T) {
		// it's important the struct fields are exported, or they will not be serialized
		person := struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}{
			"Foo",
			"Bar",
		}

		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(person)
		Is.NoErr(err)
		Is.Equal(b.String(), "{\"firstname\":\"Foo\",\"lastname\":\"Bar\"}\n")
	})
}
