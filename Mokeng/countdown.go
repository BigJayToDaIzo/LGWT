package main

import (
	"fmt"
	"io"
	"time"
)

// INJECT that interface right into that Countdown BWOI!
// Then write the struct to pass the appropriate behavior to test suite
type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

const countdownStart = 3
const finalWord = "Go!"

func Countdown(w io.Writer, s Sleeper) {
	// dependency on time.Sleep slowing test suite, MUST MOK
	// What would we name a Mock Sleep? Sleeper!
	for i := countdownStart; i > 0; i-- {
		w.Write([]byte(fmt.Sprintf("%d\n", i)))
		s.Sleep()
	}
	w.Write([]byte(fmt.Sprintf("%s\n", finalWord)))
}
