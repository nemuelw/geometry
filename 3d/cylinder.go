package spatial

import (
	"math"

	"github.com/nemzyxt/geometry/3d/shape"
)

type Cylinder struct {
	Radius, Height float64
}

func NewCylinder(radius, height float64) shape.Spatial {
	return Cylinder{Radius: radius, Height: height}
}

func (c Cylinder) Volume() float64 {
	return math.Pi * c.Radius * c.Radius * c.Height
}
