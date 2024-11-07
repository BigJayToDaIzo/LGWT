package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// TODO This is a MESS, abstraction and refcactors GALORE here
type Point struct {
	X float64
	Y float64
}

const (
	// temp configuration points set as const
	secondHandLen = 90
	minuteHandLen = 80
	hourHandLen   = 50
	x_originPoint = 150
	y_originPoint = 150

	// legit constants
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func HourHand(w io.Writer, t time.Time) {
	pt := makeHand(hourHandPoint(t), hourHandLen)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, pt.X, pt.Y)

}

func MinuteHand(w io.Writer, t time.Time) {
	pt := makeHand(minuteHandPoint(t), minuteHandLen)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, pt.X, pt.Y)
}

func SecondHand(w io.Writer, t time.Time) {
	pt := makeHand(secondHandPoint(t), secondHandLen)
	// TODO parse this from the aforementioned buffer
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, pt.X, pt.Y)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + +x_originPoint, p.Y + y_originPoint}
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

// lets do it again and make the hour hand move REAL slow
func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

// let's move the minute hand each second instead of
// on the minute marker
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
