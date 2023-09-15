package spatial

import (
	"github.com/nemzyxt/geometry/3d/shape"
)

type Cube struct {
	Side float64
}

func NewCube(side float64) shape.Spatial {
	return Cube{Side: side}
}

func (c Cube) Volume() float64 {
	return c.Side * c.Side * c.Side
}
