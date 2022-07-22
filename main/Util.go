package main

import (
	"github.com/0xrawsec/golang-utils/fsutil"
	"os"
	"syscall"

	"github.com/snnyyp/ddWin/Define"
)

func DismountOutputPartition() {
	fd := syscall.Handle(outputFileHandle.Fd())
	bytesReturned := uint32(0)
	_ = syscall.DeviceIoControl(fd, Define.FsctlDismountVolume,
		nil, 0, nil, 0, &bytesReturned, nil)
}

func GetSystemDrive() (drive string) {
	drive, _ = os.LookupEnv("SYSTEMDRIVE")
	return
}

func IsOutputFileExist() bool {
	return fsutil.IsFile(argStruct.OutputFile)
}

func EraseOrTouchOutputFile() {
	f, _ := os.Create(argStruct.OutputFile)
	f.Close()
}
