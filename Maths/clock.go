package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

// temp configuration points set as const
const secondHandLen = 90
const minuteHandLen = 80
const x_originPoint = 150
const y_originPoint = 150

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}

// TODO: refactor this to use fileio
const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func secondHand(w io.Writer, t time.Time) {
	pt := makeHand(secondHandPoint(t), secondHandLen)
	// TODO parse this from the aforementioned buffer
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, pt.X, pt.Y)
}

func SecondHand(t time.Time) Point {
	pt := secondHandPoint(t)
	//1 scale from unit circle to svg
	pt = Point{pt.X * secondHandLen, pt.Y * secondHandLen} //scale to secondHandLen

	//2 Flip triangle over the X axis to account for
	// TL Corner origin of SVG
	pt = Point{pt.X, -pt.Y} // flip the y
	//3 translate entire point to the center of the svg
	// also not in the TL corner
	pt = Point{pt.X + x_originPoint, pt.Y + y_originPoint}
	return pt
}

func minuteHand(w io.Writer, t time.Time) {
	pt := makeHand(minuteHandPoint(t), minuteHandLen)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, pt.X, pt.Y)
}

func MinuteHand(t time.Time) Point {
	pt := minuteHandPoint(t)
	//1 scale from unit circle to svg
	pt = Point{pt.X * minuteHandLen, pt.Y * minuteHandLen} //scale to minuteHandLen
	//2 Flip triangle over the X axis to account for
	// TL Corner origin of SVG
	pt = Point{pt.X, -pt.Y} // flip the y
	//3 translate entire point to the center of the svg
	// also not in the TL corner
	pt = Point{pt.X + x_originPoint, pt.Y + y_originPoint}
	return pt
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + +x_originPoint, p.Y + y_originPoint}
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return (angleToPoint(minutesInRadians(t)))
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

// let's move the minute hand each second instead of
// on the minute marker
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
