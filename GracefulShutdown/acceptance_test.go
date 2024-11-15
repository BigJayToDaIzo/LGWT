package acceptancetests

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const (
	baseBinName = "temp-testbinary"
)

func LaunchTestProgram(port string) (cleanup func(), sendInterrupt func() error, err error) {
	binName, err := buildBinary()
	if err != nil {
		return nil, nil, fmt.Errorf("could not build binary: %w", err)
	}
	sendInterrupt, kill, err := runServer(binName, port)
	cleanup = func() {
		if kill != nil {
			kill()
		}
		os.Remove(binName)
	}
	if err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("could not run server: %w", err)
	}
	return cleanup, sendInterrupt, nil
}

func buildBinary() (string, error) {
	binName := randomString(10) + "-" + baseBinName
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		return "", fmt.Errorf("could not build binary: %w", err)
	}
	return binName, nil
}

func runServer(binName, port string) (func() error, func(), error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil, fmt.Errorf("could not get current working directory: %w", err)
	}
	cmdPath := filepath.Join(dir, binName)
	cmd := exec.Command(cmdPath)
	if err := cmd.Start(); err != nil {
		return nil, nil, fmt.Errorf("failed to start server: %v", err)
	}
	kill := func() {
		_ = cmd.Process.Kill()
	}
	sendInterrupt := func() error {
		return cmd.Process.Signal(syscall.SIGTERM)
	}
	err = waitForServerListening(port)
	return sendInterrupt, kill, err
}

func waitForServerListening(port string) error {
	for i := 0; i < 30; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", port))
		if conn != nil {
			conn.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("nothing is listening on localhost: %s", port)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
