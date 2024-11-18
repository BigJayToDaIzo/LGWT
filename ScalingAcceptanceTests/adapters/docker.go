package adapters

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	startupTimeout = 5 * time.Second
	dockerfileName = "Dockerfile"
)

func StartDockerServer(
	t testing.TB,
	port string,
	binToBuild string,
) {
	t.Helper()
	ctx := context.Background()
	req := tc.ContainerRequest{
		FromDockerfile: newTCDockerfile(binToBuild),
		ExposedPorts:   []string{fmt.Sprintf("%s:%s", port, port)},
		WaitingFor:     wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(startupTimeout),
	}
	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})
}

func newTCDockerfile(binToBuild string) tc.FromDockerfile {
	return tc.FromDockerfile{
		Context:    "../../.",
		Dockerfile: dockerfileName,
		BuildArgs: map[string]*string{
			"bin_to_build": &binToBuild,
		},
		PrintBuildLog: true,
	}
}
