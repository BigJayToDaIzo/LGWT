package httpserver

import (
	"fmt"
	"net/http"

	"example.com/go-specs-greet/domain/interactions"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	// get name from request
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Greet(name))
}
