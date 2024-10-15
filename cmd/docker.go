package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/avnovoselov/snippets/cmd/docker"
	"github.com/avnovoselov/snippets/internal/helpers/exec"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "List of running containers",
	Long:  "🐳 List of running containers",
	Run: func(cmd *cobra.Command, args []string) {
		err := exec.PrintCommand("docker", "ps")
		if err != nil {
			slog.Error(err.Error())
		}
	},
}

func init() {
	dockerCmd.AddCommand(docker.AttachCmd)
}
