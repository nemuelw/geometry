package spatial

import (
	"math"

	"github.com/nemzyxt/geometry/3d/shape"
)

type Sphere struct {
	Radius float64
}

func NewSphere(radius float64) shape.Spatial {
	return Sphere{Radius: radius}
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * s.Radius * s.Radius * s.Radius) / 3
}
