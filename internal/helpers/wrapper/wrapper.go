package wrapper

import (
	"os"
	"os/exec"
)

var (
	OSGetwd = os.Getwd
	OSExit  = os.Exit

	ExecCommand = exec.Command
)
