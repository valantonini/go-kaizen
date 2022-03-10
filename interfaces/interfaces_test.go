package interfaces

import (
	"bytes"
	"github.com/matryer/is"
	"io"
	"testing"
)

type WrappedBufferedWriter struct {
	internalBuffer *bytes.Buffer
}

func (bw WrappedBufferedWriter) Write(p []byte) (n int, err error) {
	n, err = bw.internalBuffer.Write(p)
	return
}

// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

func Test_Interfaces(t *testing.T) {
	Is := is.New(t)

	t.Run("methods can accept a interface pointer when the variable's underlying type is specified correctly ", func(t *testing.T) {
		var cw io.Writer
		buffer := new(bytes.Buffer)
		cw = WrappedBufferedWriter{buffer} // new(WrappedBufferedWriter) works if no fields need to be set

		// A method that accepts an interface pointer of type *io.Writer
		HelloWorld := func(w *io.Writer) {
			(*w).Write([]byte("Hello World"))
		}

		// WrappedBufferedWriter variable cw was declared as an io.Writer so can be passed to parameters or fields
		// accepting *io.Writer when used with &
		HelloWorld(&cw)

		Is.Equal(buffer.String(), "Hello World")

		/*
			An interface value is constructed of two words of data; one word is used to point to a method table for the
			value’s underlying type, and the other word is used to point to the actual data being held by that value.

			The below will not compile: cannot use &cw (type *WrappedBufferedWriter) as type *io.Writer in argument to HelloWorld:
			*io.Writer is pointer to interface, not interface

			cw := WrappedBufferedWriter{}
			HelloWorld(&cw)
		*/
	})

	t.Run("satisfying an interface declared in a struct", func(t *testing.T) {
		type ContainerType struct {
			io.Writer
		}

		buffer := new(bytes.Buffer)

		// implementation supplied by implementation supplied on struct
		instance := ContainerType{buffer}

		// promoted from embedded io.Writer interface
		instance.Write([]byte("foo"))

		Is.Equal(buffer.String(), "foo")

		// will explode:
		// instance := ContainerType{}
		// instance.Write([]byte("foo"))

	})
}
