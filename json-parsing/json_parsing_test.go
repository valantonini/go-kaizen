package json_parsing

import (
	"encoding/json"
	"github.com/matryer/is"
	"testing"
)

type data struct {
	Id json.Number `json:"id"`
}

func Test_JsonParsing(t *testing.T) {
	Is := is.New(t)

	t.Run("json.number is a type that provides helper methods", func(t *testing.T) {
		str := `{ "id": 123 }`

		var d data
		_ = json.Unmarshal([]byte(str), &d)

		Is.Equal(json.Number("123"), d.Id)
		Is.Equal("123", d.Id.String())

		i, _ := d.Id.Int64()
		Is.Equal(int64(123), i)
	})
}
