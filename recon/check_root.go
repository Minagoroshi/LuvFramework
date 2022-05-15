package luvrecon

import "os/user"

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
