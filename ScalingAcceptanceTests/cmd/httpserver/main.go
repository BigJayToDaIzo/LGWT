package main

import (
	"net/http"

	gsg "example.com/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(gsg.Handler)
	http.ListenAndServe(":8080", handler)
}
