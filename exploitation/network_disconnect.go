package luvdisruption

import (
	"LuvFramework/utils"
)

func NetworkDisconnect() error {

	cmd := `netsh interface set interface name="Wireless Network Connection" admin=disable`

	_, err := luvutils.CmdOut(cmd)
	if err != nil {
		return err
	}

	return nil

}
