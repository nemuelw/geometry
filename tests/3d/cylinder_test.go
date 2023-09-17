package tests

import (
	"math"
	"testing"

	"github.com/nemzyxt/geometry/3d"
)

func TestNewCylinder(t *testing.T) {
	radius, height := 5.0, 10.0
	cylinder := spatial.NewCylinder(radius, height)

	if cylinder.(spatial.Cylinder).Radius != radius && cylinder.(spatial.Cylinder).Height != height {
		t.Errorf("Expected cylinder with radius %f and height %f, but got radius %f and height %f",
				float64(radius), float64(height),
				cylinder.(spatial.Cylinder).Radius, cylinder.(spatial.Cylinder).Height)		
	}
}

func TestCylinderVolume(t *testing.T) {
	radius, height := 5.0, 10.0
	cylinder := spatial.NewCylinder(radius, height)
	expectedVolume := math.Pi * radius * radius * height

	if cylinder.Volume() != expectedVolume {
		t.Errorf("Expected volume %f, but got %f", expectedVolume, cylinder.Volume())
	}
}
