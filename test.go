package main

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

func main() {
    newFd, _ := dup(os.Stdout.Fd())
    newStdout := os.NewFile(newFd, "Stdout")
    defer newStdout.Close()
    os.Stdout.Close()
    newStdout.WriteString("Hello, world!")
    os.Stdout.WriteString("Hello, world!")
}
