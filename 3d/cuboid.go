package spatial

import (
	"github.com/nemzyxt/geometry/3d/shape"
)

type Cuboid struct {
	Length, Width, Height float64
}

func NewCuboid(length, width, height float64) shape.Spatial {
	return Cuboid{Length: length, Width: width, Height: height}
}

func (c Cuboid) Volume() float64 {
	return c.Length * c.Width * c.Height
}
