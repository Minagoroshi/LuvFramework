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
	//Add a buffer to the size
	size += 1024
	return make([]byte, size), nil
}

//ArrayContains checks if an array contains a value
func ArrayContains(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
