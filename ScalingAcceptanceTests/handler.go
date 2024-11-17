package go_specs_greet

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	// get name from request
	name := request.URL.Query().Get("name")
	fmt.Fprintf(writer, "Hello, %s", name)
}
