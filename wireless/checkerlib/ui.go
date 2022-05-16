package luvcheckerlib

import (
	"errors"
	"runtime"
	"syscall"
	"unsafe"
)

//SetConsoleTitle function sets the console title of the current console window. (No support for linux)
func SetConsoleTitle(title string) error {
	var err error
	notSuppported := errors.New("Not supported on linux")

	if runtime.GOOS == "windows" {
		handle, err := syscall.LoadLibrary("Kernel32.dll")
		if err != nil {
			println(err)
		}
		defer syscall.FreeLibrary(handle)
		proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
		if err != nil {
			println(err)
		}
		_, _, err = syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	} else {
		err = notSuppported
	}
	return err
}

//StartUI function starts the UI for the program.
func StartUI() {
	if runtime.GOOS == "windows" {
		println("Starting UI")
		SetConsoleTitle("TopLib")
	} else {
		println("Starting UI")
	}
}
