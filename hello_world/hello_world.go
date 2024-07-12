package main

import "fmt"

func Hello(s string) string {
	if s == "" {
		s = "World"
	}
	return fmt.Sprintf("Hello, %s!", s)
}
