package equality

import (
	"github.com/matryer/is"
	"testing"
)

type cookingAppliance interface {
	setTemp(temp int)
}

type oven struct {
	temp int
}

// this must be a pointer, or it will change a copy
func (o *oven) setTemp(temp int) {
	o.temp = temp
}

type stove struct {
	temp int
}

// this must be a pointer, or it will change a copy
func (s *stove) setTemp(temp int) {
	s.temp = temp
}

func turnOffOven(o *oven) {
	o.setTemp(0)
}

func preheat(c cookingAppliance) {
	c.setTemp(160)
}

func Test_StructuralEquality(t *testing.T) {
	Is := is.New(t)

	t.Run("test structural equality", func(t *testing.T) {
		s := stove{}
		o := oven{}

		// either oven or stove can be passed to preheat because they both satisfy the interface despite not explicitly
		// extending it
		preheat(&s)
		Is.Equal(s.temp, 160)

		preheat(&o)
		Is.Equal(o.temp, 160)

		turnOffOven(&o)
		Is.Equal(o.temp, 0)

		// does not compile despite fields being structurally equal
		// turnOffOven(s)
	})
}
