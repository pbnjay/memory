// +build linux
// +build 386 mips mipsle arm

package memory

import "syscall"

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// This is a 32-bit system so 4GB max, but API needs uint64.
	return uint64(in.Totalram)
}
