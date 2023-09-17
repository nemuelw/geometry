package tests

import (
	"testing"

	"github.com/nemzyxt/geometry/2d"
)

const (
	side = 5.0
)

func TestNewSquare(t *testing.T) {
	square := planar.NewSquare(side)

	if square.(planar.Square).Side != side {
		t.Errorf("Expected square with side %f, but got side %f",
				side, square.(planar.Square).Side)	
	}
}

func TestSquareArea(t *testing.T) {
	square := planar.NewSquare(side)
	expectedArea := side * side
	if square.Area() != expectedArea {
		t.Errorf("Expected area of %f, but got %f", expectedArea, square.Area())
	}
}

func TestSquarePerimeter(t *testing.T) {
	square := planar.NewSquare(side)
	expectedPerimeter := 4 * side

	if square.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter of %f, but got %f", expectedPerimeter, square.Perimeter())
	}
}
