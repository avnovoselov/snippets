package git_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/git"
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func TestIsRepository_Success(t *testing.T) {
	oldExecCommand := wrapper.ExecCommand
	defer func() {
		wrapper.ExecCommand = oldExecCommand
	}()

	wrapper.ExecCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "true")
	}

	isRepository := git.IsRepository("/path/to/fake/dir")

	require.True(t, isRepository)
}

func TestIsRepository_Error(t *testing.T) {
	oldExecCommand := wrapper.ExecCommand
	defer func() {
		wrapper.ExecCommand = oldExecCommand
	}()

	wrapper.ExecCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command(":wrong/command;")
	}

	isRepository := git.IsRepository("/path/to/fake/dir")

	require.False(t, isRepository)
}

func TestIsRepository_Failed(t *testing.T) {
	oldExecCommand := wrapper.ExecCommand
	defer func() {
		wrapper.ExecCommand = oldExecCommand
	}()

	wrapper.ExecCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "not git repository")
	}

	isRepository := git.IsRepository("/path/to/fake/dir")

	require.False(t, isRepository)
}
