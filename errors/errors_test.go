package errors

import (
	"errors"
	"github.com/valantonini/go-kaizen/assert"
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
	t.Run("strongly typed errors", func(t *testing.T) {
		got := errorIfEven(2)
		want := EvenNumberError

		assert.Equal(t, got, want)
	})

	t.Run("error messages", func(t *testing.T) {
		got := errorIfEven(2).Error()
		want := "number is even"

		assert.Equal(t, got, want)
	})
}
