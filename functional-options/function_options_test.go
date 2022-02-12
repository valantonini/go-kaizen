package functional_options

import (
	"github.com/matryer/is"
	"testing"
)

func Test_FunctionalOptions(t *testing.T) {
	Is := is.New(t)

	t.Run("it defaults the host to localhost when not supplied", func(t *testing.T) {
		server, err := NewServer("")

		Is.NoErr(err)
		Is.Equal(server.Host, "localhost")
	})

	t.Run("it sets the host", func(t *testing.T) {
		server, err := NewServer("example.com")

		Is.NoErr(err)
		Is.Equal(server.Host, "example.com")
	})

	t.Run("it defaults the port to 80 when not supplied", func(t *testing.T) {
		server, err := NewServer("")

		Is.NoErr(err)
		Is.Equal(server.Port, 80)
	})

	t.Run("it sets the port", func(t *testing.T) {
		server, err := NewServer("", WithPort(8080))

		Is.NoErr(err)
		Is.Equal(server.Port, 8080)
	})

	t.Run("it errors if the port is below 0", func(t *testing.T) {
		_, err := NewServer("", WithPort(-1))

		Is.Equal(err, PortLessThanZeroError)
	})

	t.Run("it defaults the scheme to http when not supplied", func(t *testing.T) {
		server, err := NewServer("")

		Is.NoErr(err)
		Is.Equal(server.Scheme, "http")
	})

	t.Run("it sets the scheme when not supplied", func(t *testing.T) {
		server, err := NewServer("", WithScheme("https"))

		Is.NoErr(err)
		Is.Equal(server.Scheme, "https")
	})
}
