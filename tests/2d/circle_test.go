package tests

import (
	"math"
	"testing"

	"github.com/nemzyxt/geometry/2d"
)

const (
	radius = 5.0
)

func TestNewCircle(t *testing.T) {
	circle := planar.NewCircle(radius)

	if circle.(planar.Circle).Radius != radius {
		t.Errorf("Expected circle with radius %f, but got radius %f",
				radius, circle.(planar.Circle).Radius)
	}
}

func TestCircleArea(t *testing.T) {
	circle := planar.NewCircle(radius)
	expectedArea := math.Pi * radius * radius

	if circle.Area() != expectedArea {
		t.Errorf("Expected area of %f, but got %f", expectedArea, circle.Area())
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := planar.NewCircle(radius)
	expectedPerimeter := 2 * math.Pi * radius

	if circle.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter of %f, but got %f", expectedPerimeter, circle.Perimeter())
	}
}
