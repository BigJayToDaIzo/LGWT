package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"example.com/go-specs-greet/adapters"
	"example.com/go-specs-greet/adapters/httpserver"
	specs "example.com/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port    = "8080"
		baseURL = fmt.Sprintf("http://localhost:%s", port)
		driver  = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)
	adapters.StartDockerServer(t, port, "httpserver")
	specs.GreetSpecification(t, &driver)
}
