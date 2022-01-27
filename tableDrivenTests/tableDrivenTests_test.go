package tableDrivenTests

import (
	"math"
	"testing"
)

type shape interface {
	Area() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type triangle struct {
	base   float64
	height float64
}

func (c triangle) Area() float64 {
	return (c.base * c.height) * 0.5
}

func Test_TableDrivenTests(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   shape
		hasArea float64
	}{
		{name: "Rectangle", shape: rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// by wrapping each case in a t.Run you will have clearer test output on failures as it will print the name of
		// the case
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// The %#v format string will print out our struct with the values in its field
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
