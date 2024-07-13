package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
		desc  string
	}{
		{&Rectangle{10.0, 4.0}, 28.0, "rect 10x4"},
		{&Circle{10}, 10 * 2 * math.Pi, "circle 10"},
		{&Triangle{10, 3}, 23.440307, "triangle 10x3"},
	}
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if math.Abs(got-want) > 0.00001 {
			t.Errorf("got %.8g want %.8g", got, want)
		}
	}
	for _, tt := range perimeterTests {
		t.Run(tt.desc, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.want)
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
		desc  string
	}{
		{&Rectangle{10.0, 5.0}, 50.0, "rect 10x5"},
		{&Circle{10}, 10 * 10 * math.Pi, "circle 10"},
		{&Triangle{10, 3}, 15, "triangle 10x3"},
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if math.Abs(got-want) > 0.00001 {
			t.Errorf("got %.8g want %.8g", got, want)
		}
	}
	for _, tt := range areaTests {
		t.Run(tt.desc, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
