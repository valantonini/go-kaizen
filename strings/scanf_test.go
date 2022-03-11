package strings

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func Test_Scanf(t *testing.T) {
	Is := is.New(t)

	t.Run("it can parse into multiple variables", func(t *testing.T) {
		input := "10m"

		var number int
		var unit string
		fmt.Sscanf(input, "%d%s", &number, &unit)

		Is.Equal(number, 10)
		Is.Equal(unit, "m")
	})
}
