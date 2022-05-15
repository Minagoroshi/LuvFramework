package luvrecon

import (
	. "LuvFramework/utils"
	"strings"
)

func users() ([]string, error) {
	clear := []string{}
	result, err := CmdOut("net user")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(result, "\n")

	for l := range lines {
		line := lines[l]
		if !ArrayContains(line, []string{"accounts for", "------", "completed"}) {
			clear = append(clear, line)
		}
	}

	return clear, nil
}
