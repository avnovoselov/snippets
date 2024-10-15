package docker

import (
	"strings"

	"github.com/avnovoselov/snippets/internal/helpers/exec"
)

const eol = "\n"

// GetRunningContainers - get names of running container
//
//	Exec command:
//		/> docker ps --format {{.Names}}
//	and split by line break
func GetRunningContainers() ([]string, error) {
	output, err := exec.Command("docker", "ps", "--format", "{{.Names}}")

	return strings.Split(output, eol), err
}
