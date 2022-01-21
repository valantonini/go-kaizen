package errors

import (
	"errors"
	"github.com/matryer/is"
	"testing"
)

var EvenNumberError = errors.New("number is even")

func errorIfEven(num int) error {
	if num%2 == 0 {
		return EvenNumberError
	}
	return nil
}

func Test_Errors(t *testing.T) {
	is := is.New(t)
	t.Run("strongly typed errors", func(t *testing.T) {
		result := errorIfEven(2)
		is.Equal(result, EvenNumberError)
	})

	t.Run("error messages", func(t *testing.T) {
		result := errorIfEven(2).Error()
		is.Equal(result, "number is even")
	})
}
