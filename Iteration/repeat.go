package main

func Repeat(c string, times int) string {
	var s string
	for i := 0; i < times; i++ {
		s += c
	}
	return s
}
