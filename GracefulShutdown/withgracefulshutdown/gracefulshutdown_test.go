package main

import (
	"testing"
	"time"

	"github.com/quii/go-graceful-shutdown/acceptancetests"
	"github.com/quii/go-graceful-shutdown/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatalf("Failed to launch program: %v", err)
	}
	t.Cleanup(cleanup)
	// check server is up before trying to shutdown gracefully
	assert.CanGet(t, url)

	// generate request, send SIGTERM before it can complete
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	// Without gradeful shutdown, this would fail
	assert.CanGet(t, url)
	// after interrupt, the server should be shutdown, and no more requests will be accepted by the server
	assert.CantGet(t, url)
}
