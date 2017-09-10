// +build windows

package memory

import (
	"syscall"
	"unsafe"
)

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	var totalKB uint64
	kernel32, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return 0
	}
	getPhysicallyInstalledSystemMemory, err := kernel32.FindProc("GetPhysicallyInstalledSystemMemory")
	if err != nil {
		return 0
	}
	r, _, _ := getPhysicallyInstalledSystemMemory.Call(uintptr(unsafe.Pointer(&totalKB)))
	if r == 0 {
		return 0
	}
	return totalKB * 1024
}
