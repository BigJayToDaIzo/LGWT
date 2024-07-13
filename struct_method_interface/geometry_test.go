package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{10.0, 4.0}, 28.0},
		{&Circle{10}, 10 * 2 * math.Pi},
	}
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.8g want %.8g", got, want)
		}
	}
	for _, tt := range perimeterTests {
		checkPerimeter(t, tt.shape, tt.want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{10.0, 5.0}, 50.0},
		{&Circle{10}, 10 * 10 * math.Pi},
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if math.Abs(got-want) > 0.00000001 {
			t.Errorf("got %.8g want %.8g", got, want)
		}
	}
	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
