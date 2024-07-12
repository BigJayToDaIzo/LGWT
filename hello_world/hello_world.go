package main

import "fmt"

func Hello(name, l string) string {
	if name == "" {
		name = "World"
	}
	switch l {
	case "Spanish":
		return fmt.Sprintf("Hola, %s!", name)
	case "French":
		return fmt.Sprintf("Bonjour, %s!", name)
	default:
		return fmt.Sprintf("Hello, %s!", name)
	}
}
