// +build windows

package silo

import (
	"fmt"
	"os"
	"syscall"
)

// PrintConsole attaches a console and prints to it.
func PrintConsole(m string) {
	modkernel32 := syscall.NewLazyDLL("kernel32.dll")
	procAllocConsole := modkernel32.NewProc("AllocConsole")
	r0, _, _ := syscall.Syscall(procAllocConsole.Addr(), 0, 0, 0, 0)
	if r0 == 0 { // allocation failed, probably process already has a console
		os.Exit(1)
	}
	out, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		os.Exit(1)
	}
	os.Stdout = os.NewFile(uintptr(out), "/dev/stdout")

	fmt.Printf("%s\n", m)
}
