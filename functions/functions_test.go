package functions

import (
	"github.com/matryer/is"
	"testing"
)

type coord struct {
	x int
	y int
}

// namedReturns demonstrates named return variables in the signature. Not always the most obvious
func namedReturns(c coord) (x, y int) {
	// can be directly assigned to
	x = c.x
	y = c.y

	// naked return
	return
}

func Test_Functions(t *testing.T) {
	Is := is.New(t)

	x, y := namedReturns(coord{3, 4})
	Is.Equal(x, 3)
	Is.Equal(y, 4)
}
