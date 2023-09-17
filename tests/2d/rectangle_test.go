package tests

import (
	"testing"

	"github.com/nemzyxt/geometry/2d"
)

const (
	length = 5.0
	width = 4.0
)

func TestNewRectangle(t *testing.T) {
	rectangle := planar.NewRectangle(length, width)

	if rectangle.(planar.Rectangle).Length != length && rectangle.(planar.Rectangle).Width != width {
		t.Errorf("Expected rectangle with length %f and width %f, but got length %f and width %f",
				length, width, rectangle.(planar.Rectangle).Length, rectangle.(planar.Rectangle).Width)	
	}
}

func TestRectangleArea(t *testing.T) {
	rectangle := planar.NewRectangle(length, width)
	expectedArea := length * width

	if rectangle.Area() != expectedArea {
		t.Errorf("Expected area of %f, but got %f", expectedArea, rectangle.Area())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rectangle := planar.NewRectangle(length, width)
	expectedPerimeter := 2 * (length + width)

	if rectangle.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter of %f, but got %f", expectedPerimeter, rectangle.Perimeter())
	}
}
