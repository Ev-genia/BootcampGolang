package main

import (
	"errors"
	"os"
	"os/exec"
)

func main() {
	var (
		cmdStr string
		args   []string
	)
	if len(os.Args) == 1 {
		cmdStr = "echo"
	} else {
		cmdStr = os.Args[1]
		args = os.Args[2:]
	}
	cmd := exec.Command(cmdStr, args...)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
