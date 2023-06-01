package main

import (
	"bufio"
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
	pathScanner := bufio.NewScanner(os.Stdin)
	for pathScanner.Scan() {
		// line, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Fatal("Error reading: ", err)
		// }
		// line = strings.TrimSpace(line)
		// if len(line) == 0 {
		// 	break
		// }
		line := pathScanner.Text()
		if len(line) == 0 {
			break
		}
		args = append(args, line)
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
