package main

import (
	"os"
	"time"
)

func main() {
	// stdout for main
	sleeper := &ConfigurableSleeper{1 * time.Millisecond, time.Sleep}
	Countdown(os.Stdout, sleeper)

	// configurable sleeper
	csleeper := &ConfigurableSleeper{333 * time.Millisecond, time.Sleep}
	Countdown(os.Stdout, csleeper)

	// http module next? Countdown FOR TEH WEBZ?
}
