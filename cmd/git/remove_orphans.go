package git

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/avnovoselov/snippets/internal/helpers/exec"
	"github.com/avnovoselov/snippets/internal/helpers/git"
	"github.com/avnovoselov/snippets/internal/helpers/path"
)

var RemoveOrphansCmd = &cobra.Command{
	Use:     "remove-orphans",
	Aliases: []string{"ro"},
	Short:   "🗑️ Remove orphans local branch",
	Long: `🗑️ Remove orphans local branch.
Get all local branches without a remote and remove them.  
`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.Command("bash", "-c", "git fetch -p && for branch in $(git for-each-ref --format '%(refname) %(upstream:track)' refs/heads | awk '$2 == \"[gone]\" {sub(\"refs/heads/\", \"\", $1); print $1}'); do git branch -D $branch; done\n")
		if err != nil {
			slog.Error(err.Error())
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		git.MustRepository(path.MustCurrentPath())
	},
}
