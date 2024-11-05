package main

import (
	"os"
	"time"

	c "example.com/clockface"
)

// const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
// <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
// <svg xmlns="http://www.w3.org/2000/svg"
//      width="100%"
//      height="100%"
//      viewBox="0 0 300 300"
//      version="2.0">`
//
// const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
//
// const svgEnd = `</svg>`

func main() {
	t := time.Now()
	c.SVGWriter(os.Stdout, t)

}

// func secondHandTag(p c.Point) string {
// 	return fmt.Sprintf(`<line x1="150" y1="150" x2="%v" y2="%v" style="fill:none;stroke:black;stroke-width:3px;"/>`, p.X, p.Y)
// }
