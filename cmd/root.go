package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snippets",
	Short: "Little helper for automation of routine processes",
	Long:  `🚀 Little helper for automation of routine processes`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			slog.Error(err.Error())
		}
	},
}

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "--verbose")

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	rootCmd.AddCommand(gitCmd, dockerCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
