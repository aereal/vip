package main

import (
	"os"
	"os/exec"
)

var Runner = func(cmd *exec.Cmd) error {
	return cmd.Run()
}

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return Runner(cmd)
}
