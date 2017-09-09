// +build windows

package memory

import (
	"syscall"
	"unsafe"
)

var (
	kernel32                           = syscall.MustLoadDLL("kernel32.dll")
	getPhysicallyInstalledSystemMemory = kernel32.MustFindProc("GetPhysicallyInstalledSystemMemory")
)

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	var totalKB uint64
	r, _, _ := getPhysicallyInstalledSystemMemory.Call(uintptr(unsafe.Pointer(&totalKB)))
	if r == 0 {
		return 0
	}
	return totalKB * 1024
}
