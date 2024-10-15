package wrapper

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	OSGetwd = os.Getwd
	OSExit  = os.Exit

	ExecCommand = exec.Command

	FMTPrintln = fmt.Println
)
