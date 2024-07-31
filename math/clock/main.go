package main

import (
	"os"
	"time"

	"example.com/clock"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
