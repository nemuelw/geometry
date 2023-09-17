package tests

import (
	"testing"

	"github.com/nemzyxt/geometry/3d"
)

const (
	side = 5
)

func TestNewCube(t *testing.T) {
	cube := spatial.NewCube(side)

	if cube.(spatial.Cube).Side != side {
		t.Errorf("Expected cube with side %f, but got side %f",
				float64(side), cube.(spatial.Cube).Side)		
	}
}

func TestCubeVolume(t *testing.T) {
	cube := spatial.NewCube(side)
	expectedVolume := float64(side * side * side)
	if cube.Volume() != expectedVolume {
		t.Errorf("Expected volume %f, but got %f", expectedVolume, cube.Volume())
	}
}
