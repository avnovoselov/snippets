package exec_test

import (
	"testing"

	"github.com/avnovoselov/snippets/internal/helpers/exec"
	"github.com/stretchr/testify/require"
)

func TestCommand_Success(t *testing.T) {
	output, err := exec.Command("pwd")

	require.NotEmpty(t, output)
	require.NotEqual(t, output[len(output)-2], "\n")
	require.NoError(t, err)
}
