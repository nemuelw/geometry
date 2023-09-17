package planar

import (
	"math"

	"github.com/nemzyxt/geometry/2d/shape"
)

type Triangle struct {
	SideA, SideB, SideC float64
}

func NewTriangle(side_a, side_b, side_c float64) shape.Planar {
	return Triangle{SideA: side_a, SideB: side_b, SideC: side_c}
}

func (t Triangle) Area() float64 {
	s := 0.5 * (t.SideA + t.SideB + t.SideC)
	return math.Sqrt((s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC)))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
