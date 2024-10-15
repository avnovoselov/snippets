package docker_test

import (
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/docker"
)

func TestRunBashInContainer(t *testing.T) {
	oldExecCommand := wrapper.ExecCommand
	defer func() {
		wrapper.ExecCommand = oldExecCommand
	}()

	containerName := "fake_container_name"
	wrapper.ExecCommand = func(name string, arg ...string) *exec.Cmd {
		require.Contains(t, arg, containerName)

		return &exec.Cmd{}
	}

	err := docker.RunBashInContainer(containerName)
	require.Error(t, err)
}
