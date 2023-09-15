package planar

import (
	"math"

	"github.com/nemzyxt/geometry/2d/shape"
)

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) shape.Planar {
	return Circle{Radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
