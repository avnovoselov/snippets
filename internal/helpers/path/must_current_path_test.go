package path_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/path"
	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func TestMustCurrentPath_Success(t *testing.T) {
	oldOSExit := wrapper.OSExit
	defer func() {
		wrapper.OSExit = oldOSExit
	}()

	wrapper.OSExit = func(int) {
		require.True(t, false) // not executed
	}

	path.MustCurrentPath()
}

func TestMustCurrentPath_Failed(t *testing.T) {
	oldOSExit, oldOSGetwd := wrapper.OSExit, wrapper.OSGetwd
	defer func() {
		wrapper.OSExit = oldOSExit
		wrapper.OSGetwd = oldOSGetwd
	}()

	wrapper.OSExit = func(code int) { require.Equal(t, 1, code) }
	wrapper.OSGetwd = func() (string, error) { return "", io.EOF }

	path.MustCurrentPath()
}
