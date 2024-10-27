package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (string, error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

// struct{} is the lightest aka empty type we can pass
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// closing the channel as complete is all the intel we need!
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
