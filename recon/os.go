package luvrecon

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

//GetWindowsVersion returns the version of windows running on the host machine
func GetWindowsVersion() (int, error) {
	cmd := exec.Command("cmd", "ver")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	osStr := strings.Replace(out.String(), "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	tmp1 := strings.Index(osStr, "[Version")
	tmp2 := strings.Index(osStr, "]")
	if tmp1 == -1 || tmp2 == -1 {
		return 0, errors.New("version string has wrong format")
	}
	longVer := osStr[tmp1+9 : tmp2]
	majorVerStr := strings.SplitN(longVer, ".", 2)[0]
	majorVerInt, err := strconv.Atoi(majorVerStr)
	if err != nil {
		return 0, errors.New("version could not be converted to int")
	}
	return majorVerInt, nil
}
