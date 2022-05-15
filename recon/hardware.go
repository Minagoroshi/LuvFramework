package luvrecon

import (
	luvutils "LuvFramework/utils"
	"golang.org/x/sys/windows/registry"
	"strings"
)

//MachineGuid returns the machine guid
func MachineGuid() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue("MachineGuid")
	if err != nil {
		return "", err
	}
	return s, nil
}

//GpuName returns the GPU info for the machine
func GpuName() ([]string, error) {
	gpuName := []string{}
	out, err := luvutils.CmdOut("wmic path win32_VideoController get name")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if !luvutils.ArrayContains(line, []string{"Name"}) {
			gpuName = append(gpuName, line)
		}
	}

	return gpuName, nil
}
