package closures

import (
	"github.com/matryer/is"
	"testing"
)

var makeCounterClosure = func() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

var makeParamCounterClosure = func(start int) func() int {
	return func() int {
		return start + 1
	}
}

func TestClosures(t *testing.T) {
	Is := is.New(t)

	t.Run("local vars are stored in higher order functions closure", func(t *testing.T) {
		counter := makeCounterClosure()

		Is.Equal(counter(), 1)
		Is.Equal(counter(), 2)
	})

	t.Run("local vars are stored in higher order functions closure are independent", func(t *testing.T) {
		counter1 := makeCounterClosure()
		counter2 := makeCounterClosure()

		Is.Equal(counter1(), 1)
		Is.Equal(counter2(), 1)
		Is.Equal(counter1(), 2)
		Is.Equal(counter2(), 2)
	})

	t.Run("variables are passed as values", func(t *testing.T) {
		start := 0
		counter := makeParamCounterClosure(start)

		Is.Equal(counter(), 1)
		Is.Equal(counter(), 1)
	})
}
