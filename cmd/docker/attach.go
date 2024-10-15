package docker

import (
	"log/slog"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/avnovoselov/snippets/internal/helpers/docker"
)

var AttachCmd = &cobra.Command{
	Use:   "attach",
	Short: "List of running containers",
	Long:  "🐳 List of running containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers, err := docker.GetRunningContainers()
		if err != nil {
			slog.Error(err.Error())
			return
		}

		container, err := selectContainer(containers)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		err = docker.RunBashInContainer(container)
		if err != nil {
			slog.Error(err.Error())
		}
	},
}

func selectContainer(containers []string) (string, error) {
	prompt := promptui.Select{
		Label:             "Select container",
		Items:             containers,
		Size:              10,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			c := strings.Replace(strings.ToLower(containers[index]), " ", "", -1)
			i := strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(c, i)
		},
	}
	_, container, err := prompt.Run()

	return container, err
}
