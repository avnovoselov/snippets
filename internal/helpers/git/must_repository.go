package git

import (
	"fmt"
	"log/slog"

	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func MustRepository(path string) {
	if !IsRepository(path) {
		slog.Error(fmt.Sprintf("Path \"%s\" should be a git repository", path))
		wrapper.OSExit(1)
	}
}
