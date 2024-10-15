package docker

import (
	"strings"

	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

const eol = "\n"

// GetRunningContainers - get names of running container
//
//	Exec command:
//		/> docker ps --format {{.Names}}
//	and split by line break
func GetRunningContainers() ([]string, error) {
	cmd := wrapper.ExecCommand("docker", "ps", "--format", "{{.Names}}")
	output, err := cmd.Output()

	return strings.Split(string(output), eol), err
}
