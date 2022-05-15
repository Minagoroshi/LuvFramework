package luvdisruption

import "os"

//The Wipe function is used to wipe the filesystem of a given path.
func Wipe(path string) {

	//Wipe the drive
	err := os.RemoveAll(path)
	if err != nil {
		return
	}

}
