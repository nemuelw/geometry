package planar

import (
	"github.com/nemzyxt/geometry/2d/shape"
)

type Rectangle struct {
	Length float64
	Width float64
}

func NewRectangle(length, width float64) shape.Planar {
	return Rectangle{Length: length, Width: width}
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}
