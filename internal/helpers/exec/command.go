package exec

import (
	"os"
	"strings"

	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func Command(name string, arg ...string) (string, error) {
	cmd := wrapper.ExecCommand(name, arg...)
	out, err := cmd.Output()

	return strings.Trim(string(out), "\n\t"), err
}

func PrintCommand(name string, arg ...string) error {
	output, err := Command(name, arg...)
	if err != nil {
		return err
	}

	_, _ = wrapper.FMTPrintln(output)

	return nil
}

func CommandInteractive(name string, arg ...string) error {
	cmd := wrapper.ExecCommand(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
