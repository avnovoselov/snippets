package docker

import "github.com/avnovoselov/snippets/internal/helpers/exec"

// RunBashInContainer - run bash in container
//
//	Execute command:
//		/> docker exec -ti <container> bash
func RunBashInContainer(container string) error {
	return exec.CommandInteractive("docker", "exec", "-ti", container, "bash")
}
