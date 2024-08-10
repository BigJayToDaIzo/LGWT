package main_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	go_specs_greet "example.com/go-specs-greet"
	"example.com/go-specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			// set to false for less spamzes
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8000/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8000"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8008", Client: &client}
	specifications.GreetSpecification(t, driver)
}
