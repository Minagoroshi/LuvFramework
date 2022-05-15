package luvutils

import (
	"os/exec"
)

//The CmdOut function is used to execute a command
func CmdOut(command string) (string, error) {
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

//The AllocMem function is used to allocate memory
func AllocMem(size int) ([]byte, error) {
	return make([]byte, size), nil
}
