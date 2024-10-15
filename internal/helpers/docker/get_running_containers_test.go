package docker_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/docker"
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func TestGetRunningContainers_Success(t *testing.T) {
	oldExecCommand := wrapper.ExecCommand
	defer func() {
		wrapper.ExecCommand = oldExecCommand
	}()

	wrapper.ExecCommand = func(string, ...string) *exec.Cmd {
		return exec.Command("echo", "first\nsecond\nthird")
	}

	containers, err := docker.GetRunningContainers()
	require.NoError(t, err)
	require.Contains(t, containers, "first")
	require.Contains(t, containers, "second")
	require.Contains(t, containers, "third")
}
