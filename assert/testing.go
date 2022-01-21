package assert

import (
	"testing"
)

// Equal compares 2 inputs and errors if not equal
func Equal(t testing.TB, got, want interface{}) {
	// t.Helper() ensure error is reported at callsite
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
