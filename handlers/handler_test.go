package handlers

import (
	"io"
	"testing"
)

type handler func(w io.Writer)

func (h *handler) Handle(w io.Writer) {
	h.Handle(w)
}

func Test_Handlers(t *testing.T) {
}
