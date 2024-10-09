package exec

import (
	"strings"

	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func Command(name string, arg ...string) (string, error) {
	cmd := wrapper.ExecCommand(name, arg...)
	out, err := cmd.Output()

	return strings.Trim(string(out), "\n\t"), err
}
