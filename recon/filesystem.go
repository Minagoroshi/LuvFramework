package luvrecon

import (
	"io/ioutil"
	"os"
)

//The GetDrives function returns a list of drives on the system.
func GetDrives() ([]string, error) {
	var drives []string

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

//GetFileData returns the data of a file as a byte slice
func GetFileData(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetFileSize returns the size of a file in bytes
func GetFileSize(file string) (int64, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
