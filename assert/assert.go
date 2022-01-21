package assert

import "testing"

func Equal(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func NotNil(t testing.TB, got interface{}) {
	t.Helper()
	if got == nil {
		t.Error("got was nil")
	}
}

func Nil(t testing.TB, got interface{}) {
	t.Helper()
	if got: != nil {
		t.Error("got was not nil")
	}
}
