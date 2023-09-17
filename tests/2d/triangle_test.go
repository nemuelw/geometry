package tests

import (
	"math"
	"testing"

	"github.com/nemzyxt/geometry/2d"
)

const (
	a = 3; b = 4; c = 5
)

func TestNewTriangle(t *testing.T) {
	triangle := planar.NewTriangle(a, b, c)

	if triangle.(planar.Triangle).SideA != a {
		t.Errorf("Expected triangle with side a %f, but got side a %f",
				float64(a), triangle.(planar.Triangle).SideA)	
	}
}

func TestTriangleArea(t *testing.T) {
	triangle := planar.NewTriangle(a, b, c)
	s := 0.5 * (a + b + c)
	expectedArea := math.Sqrt((s * (s - a) * (s - b) * (s - c)))
	if triangle.Area() != expectedArea {
		t.Errorf("Expected area of %f, but got %f", expectedArea, triangle.Area())
	}
}

func TestTrianglePerimeter(t *testing.T) {
	triangle := planar.NewTriangle(a, b, c)
	expectedPerimeter := float64(a + b + c)

	if triangle.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter of %f, but got %f", expectedPerimeter, triangle.Perimeter())
	}
}
