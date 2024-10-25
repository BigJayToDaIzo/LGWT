package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// INJECT that interface right into that Countdown BWOI!
// Then write the struct to pass the appropriate behavior to test suite
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	// dependency on time.Sleep slowing test suite, MUST MOK
	// What would we name a Mock Sleep? Sleeper!
	for i := 3; i > 0; i-- {
		w.Write([]byte(fmt.Sprintf("%d\n", i)))
		s.Sleep()
	}
	w.Write([]byte("Go!"))
}

func main() {
	// stdout for main
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
	// http module next? Countdown FOR TEH WEBZ?
	// wooo, COUNTDOWN FOR THE
}
