package initFunction

import (
	"github.com/matryer/is"
	"testing"
)

// https://golangdocs.com/init-function-in-golang
func Test_InitFunction(t *testing.T) {
	Is := is.New(t)

	// init functions are special functions function executes after the package is imported and maintains the order of
	// execution. That means multiple init functions can be defined in a file, and they will be called one after another
	// maintaining the order. Can be a dark pattern as it's not obvious. File order is alphabetical order and should not
	// be relied upon. File can contain multiple init(). will be executed in order.
	t.Run("executes on module import", func(t *testing.T) {
		Is.Equal(SetByInit, "set by init")
		Is.Equal(SetByInit2, "set by init2")
	})
}
