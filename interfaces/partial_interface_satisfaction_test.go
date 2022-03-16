package interfaces

import (
	"testing"

	"github.com/matryer/is"
)

// Fooer has 2 methods that need to be satisfied
type Fooer interface {
	GetFoo() string
	GetBaz() string
}

type Barer interface {
	GetBar() string
}

// FooBarer is a struct that should implement the Fooer and Barer interfaces via embeds
type FooBarer struct {
	Fooer
	Barer
}

func GetFoo(fb FooBarer) string {
	return fb.GetFoo()
}

type FooerImpl struct {
	Fooer // this embed will ensure FooerImple satisfies the Fooer interface
}

// GetFoo is the only method called in execution. This receiver is called instead of the embed's
func (f FooerImpl) GetFoo() string {
	return "foo"
}

type PartialImpl struct {
	FooBarer
}

func Test_PartialInterfaceSatisfaction(t *testing.T) {
	Is := is.New(t)

	partialImpl := FooBarer{FooerImpl{}, nil}

	Is.Equal(partialImpl.GetFoo(), "foo")
}
