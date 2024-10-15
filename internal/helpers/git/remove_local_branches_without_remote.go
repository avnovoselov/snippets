package git

import (
	"github.com/avnovoselov/snippets/internal/helpers/exec"
)

// RemoveLocalBranchesWithoutRemote
//
//	Execute command:
//		/> git fetch -p && for branch in $(git for-each-ref --format '%(refname) %(upstream:track)' refs/heads | awk '$2 == "[gone]" {sub("refs/heads/", "", $1); print $1}'); do git branch -D $branch; done\n
func RemoveLocalBranchesWithoutRemote() error {
	return exec.PrintCommand("bash", "-c", "git fetch -p && for branch in $(git for-each-ref --format '%(refname) %(upstream:track)' refs/heads | awk '$2 == \"[gone]\" {sub(\"refs/heads/\", \"\", $1); print $1}'); do git branch -D $branch; done\n")
}
