package main

import (
	"net/http"

	"example.com/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	_ = http.ListenAndServe(":8080", handler)
}
