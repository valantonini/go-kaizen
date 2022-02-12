package functional_options

import (
	"github.com/matryer/is"
	"testing"
)

func Test_FunctionalOptions(t *testing.T) {
	Is := is.New(t)

	t.Run("it defaults the host to local host when not supplied", func(t *testing.T) {
		server, err := NewServer("")

		Is.NoErr(err)
		Is.Equal(server.Host, "localhost")
	})

	t.Run("it sets the host", func(t *testing.T) {
		server, err := NewServer("example.com")

		Is.NoErr(err)
		Is.Equal(server.Host, "example.com")
	})
}
