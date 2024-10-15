package git_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/git"
	"github.com/avnovoselov/snippets/internal/helpers/test"
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func TestRemoveLocalBranchesWithoutRemote(t *testing.T) {
	mock := func(name string, arg ...string) *exec.Cmd {
		return &exec.Cmd{}
	}

	restore := test.KeepOLD[func(string, ...string) *exec.Cmd](&wrapper.ExecCommand, &mock)
	defer restore()

	err := git.RemoveLocalBranchesWithoutRemote()
	require.Error(t, err)
}
