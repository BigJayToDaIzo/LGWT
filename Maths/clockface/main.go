package main

import (
	"os"
	"time"

	s "example.com/clockface/svg"
)

func main() {
	t := time.Now()
	s.SVGWriter(os.Stdout, t)
}
