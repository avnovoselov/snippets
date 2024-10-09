package exec

import (
	"os/exec"
	"strings"
)

func Command(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.Output()

	return strings.Trim(string(out), "\n\t"), err
}
