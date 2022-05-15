package luvrecon

import (
	"io/ioutil"
	"os"
)

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

//ListFiles lists all files in a directory and returns them as a slice of strings
func ListFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}
