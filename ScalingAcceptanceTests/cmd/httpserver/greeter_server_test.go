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
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)
	adapters.StartDockerServer(t, port, dockerFilePath)
	specs.GreetSpecification(t, driver)
}
