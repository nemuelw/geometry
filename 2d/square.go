package planar

import (
	"github.com/nemzyxt/geometry/2d/shape"
)

type Square struct {
	Side float64
}

func NewSquare(side float64) shape.Planar {
	return Square{Side: side}
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}
