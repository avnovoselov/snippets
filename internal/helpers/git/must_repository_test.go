package git_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/git"
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func TestMustRepository_Success(t *testing.T) {
	oldOSExit := wrapper.OSExit
	defer func() {
		wrapper.OSExit = oldOSExit
	}()

	wrapper.OSExit = func(int) {
		require.True(t, false) // not executed
	}

	wrapper.ExecCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "true")
	}

	git.MustRepository("/path/to/fake/dir")
}

func TestMustRepository_Failed(t *testing.T) {
	oldOSExit := wrapper.OSExit
	defer func() {
		wrapper.OSExit = oldOSExit
	}()

	wrapper.OSExit = func(code int) {
		require.Equal(t, 1, code) // not executed
	}

	wrapper.ExecCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "false")
	}

	git.MustRepository("/path/to/fake/dir")
}
