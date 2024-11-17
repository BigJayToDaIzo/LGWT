package httpserver

import (
	"fmt"
	"net/http"

	gsg "example.com/go-specs-greet"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	// get name from request
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, gsg.Greet(name))
}
