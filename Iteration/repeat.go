package main

// Repeat returns a string with a character repeated n times
func Repeat(c string, times int) string {
	var s string
	for i := 0; i < times; i++ {
		s += c
	}
	return s
}
