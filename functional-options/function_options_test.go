package functional_options

import (
	"github.com/matryer/is"
	"testing"
)

func Test_FunctionalOptions(t *testing.T) {
	Is := is.New(t)

	t.Run("it sets the host", func(t *testing.T) {
		server, err := NewServer("example.com")

		Is.NoErr(err)
		Is.Equal(server.Host, "example.com")
	})
}
