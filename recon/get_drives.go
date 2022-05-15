package luvrecon

import "os"

//The GetDrives function returns a list of drives on the system.
func GetDrives() ([]string, error) {
	drives := []string{}

	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			drives = append(drives, string(drive)+":\\")
			err := f.Close()
			if err != nil {
				return nil, err
			}
		}
	}
	return drives, nil
}
