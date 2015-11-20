package main

import (
	"os"
	"os/exec"
	"strings"
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

func capture(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(out)), nil
	}
}
