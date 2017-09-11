// +build linux

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
	// If this is a 32-bit system, then this is uint32 (max 4GB)
	// So we convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}
