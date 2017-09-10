// +build darwin
// +build 386 arm

package memory

import "syscall"

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	s, err := syscall.SysctlUint32("hw.memsize")
	if err != nil {
		return 0
	}
	// This is a 32-bit system so 4GB max, but API needs uint64.
	return uint64(s)
}
