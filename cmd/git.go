package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/avnovoselov/snippets/cmd/git"
	"github.com/avnovoselov/snippets/internal/helpers/path"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Check is current directory a git repository",
	Long:  `✅ Check is current directory a git repository`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info(fmt.Sprintf("✅  current directory is a git repository"))
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		path.MustCurrentPath()
	},
}

func init() {
	gitCmd.AddCommand(git.RemoveOrphansCmd)
}
