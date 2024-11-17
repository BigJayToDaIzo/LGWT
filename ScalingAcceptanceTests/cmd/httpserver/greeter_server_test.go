package main_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	gsg "example.com/go-specs-greet"
	specs "example.com/go-specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()
	req := tc.ContainerRequest{
		FromDockerfile: tc.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			// set to false after stable test suite
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
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
	driver := gsg.Driver{BaseURL: "http://localhost:8080", Client: &client}
	specs.GreetSpecification(t, driver)
}
