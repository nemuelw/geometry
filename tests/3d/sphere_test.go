package tests

import (
	"math"
	"testing"

	"github.com/nemzyxt/geometry/3d"
)

const (
	radius = 5
)

func TestNewSphere(t *testing.T) {
	sphere := spatial.NewSphere(radius)

	if sphere.(spatial.Sphere).Radius != radius {
		t.Errorf("Expected sphere with radius %f, but got radius %f",
				float64(radius), sphere.(spatial.Sphere).Radius)		
	}
}

func TestSphereVolume(t *testing.T) {
	sphere := spatial.NewSphere(radius)
	expectedVolume := (4 * math.Pi * radius * radius * radius) / 3

	if sphere.Volume() != expectedVolume {
		t.Errorf("Expected volume %f, but got %f", expectedVolume, sphere.Volume())
	}
}
