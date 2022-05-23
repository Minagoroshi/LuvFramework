package luvrecon

import (
	. "github.com/Minagoroshi/LuvFramework/utils"
	"os/user"
	"strings"
)

// CheckRoot Checks if the current user is root
func CheckRoot() bool {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	if usr.Username != "root" {
		return false
	}
	return true
}

func GetUsers() ([]string, error) {
	var clear []string
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
