package exec_test

import (
	osExec "os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/exec"
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func TestCommand_Success(t *testing.T) {
	output, err := exec.Command("pwd")

	require.NotEmpty(t, output)
	require.NotEqual(t, output[len(output)-2], "\n")
	require.NoError(t, err)
}

func TestPrintCommand_Success(t *testing.T) {
	err := exec.PrintCommand("pwd")

	require.NoError(t, err)
}

func TestPrintCommand_Failed(t *testing.T) {
	oldExecCommand := wrapper.ExecCommand
	defer func() {
		wrapper.ExecCommand = oldExecCommand
	}()

	wrapper.ExecCommand = func(string, ...string) *osExec.Cmd {
		return &osExec.Cmd{}
	}

	err := exec.PrintCommand("pwd")

	require.Error(t, err)
}

func TestCommandInteractive_Success(t *testing.T) {
	err := exec.CommandInteractive("pwd")

	require.NoError(t, err)
}
