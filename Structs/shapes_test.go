package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 62.83185307179586},
	}
	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt, got, tt.want)
			}

		})
	}
	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := r.Perimeter()
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		c := Circle{10}
		got := c.Perimeter()
		want := 62.83185307179586
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt, got, tt.want)
			}

		})
	}
	// checkArea := func(t *testing.T, s Shape, want float64) {
	// 	t.Helper()
	// 	got := s.Area()
	// 	if got != want {
	// 		t.Errorf("%#v got %g want %g", s, got, want)
	// 	}
	// }
}
