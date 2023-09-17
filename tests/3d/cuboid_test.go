package tests

import (
	"testing"

	"github.com/nemzyxt/geometry/3d"
)

const (
	length = 1; width = 2; height = 3
)

func TestNewCuboid(t *testing.T) {
	cuboid := spatial.NewCuboid(length, width, height)

	if cuboid.(spatial.Cuboid).Length != length {
		t.Errorf("Expected cuboid with length %f, but got length %f",
				float64(length), cuboid.(spatial.Cuboid).Length)		
	}
}

func TestCuboidVolume(t *testing.T) {
	cuboid := spatial.NewCuboid(length, width, height)
	expectedVolume := float64(length * width * height)

	if cuboid.Volume() != expectedVolume {
		t.Errorf("Expected volume %f, but got %f", expectedVolume, cuboid.Volume())
	}
}
