package main

import "fmt"

func Hello(name, l string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(getPrefix(l), name)
}
func getPrefix(l string) string {
	switch l {
	case "Spanish":
		return "Hola, %s!"
	case "French":
		return "Bonjour, %s!"
	default:
		return "Hello, %s!"
	}
}
