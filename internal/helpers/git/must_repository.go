package git

import (
	"fmt"
	"log/slog"
	"os"
)

func MustRepository(path string) {
	if !IsRepository(path) {
		slog.Error(fmt.Sprintf("Path \"%s\" should by a git repository", path))
		os.Exit(1)
	}
}
