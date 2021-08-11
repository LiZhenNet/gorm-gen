package core

import (
	"os/exec"
)

func RunCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func ProcessGoFile(dir string) error {
	if dir == "" {
		dir = "."
	}
	err := RunCommand(dir, "goimports", "-w", ".")
	if err != nil {
		return err
	}
	err = RunCommand(dir, "gofmt", "-w", ".")
	if err != nil {
		return err
	}
	return nil
}
