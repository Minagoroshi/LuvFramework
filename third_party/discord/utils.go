package luvdiscord

import (
	"encoding/base64"
)

//func PartyInvitation() {
//
//	if antivm.IsVirtualDisk() || antimem.ByMemWatcher() || antidebug.ByProcessWatcher() || antivm.ByMacAddress() {
//		os.Exit(0)
//	}
//}

func cake(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func icing(str string) string {
	data := base64.StdEncoding.EncodeToString([]byte(str))
	return string(data)
}
