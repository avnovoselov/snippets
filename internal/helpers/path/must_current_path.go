package path

import (
	"log/slog"
	"os"
)

func MustCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return path
}
