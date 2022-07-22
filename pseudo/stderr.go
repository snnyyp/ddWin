package Pseudo

import (
	"os"
	"syscall"
)

func dup(fd uintptr) (uintptr, error) {
	currentProcessHandle, _ := syscall.GetCurrentProcess()
	var newFd syscall.Handle

	err := syscall.DuplicateHandle(
		currentProcessHandle,
		syscall.Handle(fd),
		currentProcessHandle,
		&newFd,
		0,
		true,
		syscall.DUPLICATE_SAME_ACCESS,
	)

	if err != nil {
		return 0, err
	}
	return uintptr(newFd), nil
}

type Stderr struct {
	Stdout
}

func (s *Stderr) New() {
	originalFd := os.Stderr.Fd()
	newFd, _ := dup(originalFd)
	s.file = os.NewFile(newFd, "Stderr")
}
