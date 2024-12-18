package clockface

import (
	"math"
	"testing"
	"time"
)

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(12, 0, 0), Point{0, 1}},
		{simpleTime(15, 0, 0), Point{1, 0}},
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v, but got %v", c.point, got)
			}
		})
	}

}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v, but got %v", c.point, got)
			}
		})
	}
}

// len of line
func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		// svg weirdness -y is UP
		// origin is top left,
		// so y increases as it goes DOWN
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v, but got %v", c.point, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(3, 0, 0), math.Pi / 2},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(9, 0, 0), math.Pi * 1.5},
		{simpleTime(12, 0, 0), 0},
		{simpleTime(15, 0, 0), math.Pi / 2},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(24, 0, 0), 0}, // not real just pushing edges
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
		// TODO understand how to press this test much further with that strange mathgix
		// {simpleTime(0, 1, 45), math.Pi / ((6 * 60 * 60) / 105)},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 15, 0), math.Pi / 2},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), math.Pi * 1.5},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

// angle of line
func TestSecondsInRadians(t *testing.T) {
	// C = 2πr so 30s = π radians
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 15), math.Pi / 2},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), math.Pi * 1.5},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if !roughlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func simpleTime(hour, min, sec int) time.Time {
	return time.Date(312, time.October, 28, hour, min, sec, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThresh = 1.e-7
	return math.Abs(a-b) < equalityThresh
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
