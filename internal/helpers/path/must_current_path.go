package path

import (
	"log/slog"

	"github.com/avnovoselov/snippets/internal/helpers/wrapper"
)

func MustCurrentPath() string {
	path, err := wrapper.OSGetwd()
	if err != nil {
		slog.Error(err.Error())
		wrapper.OSExit(1)
	}

	return path
}
