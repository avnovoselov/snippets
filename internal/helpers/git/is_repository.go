package git

import (
	"fmt"
	"log/slog"

	"github.com/avnovoselov/snippets/internal/helpers/exec"
)

// IsRepository
//
//	Execute command:
//	/> git --git-dir <path> rev-parse --is-inside-work-tree
//	and check output equals "true"
func IsRepository(path string) bool {
	out, err := exec.Command("git", "--git-dir", fmt.Sprintf("%s/.git", path), "rev-parse", "--is-inside-work-tree")
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	return string(out) == "true"
}
