package memory

import "syscall"

// TotalMemory returns the total system memory in bytes, or 0 if
// available memory could not be determined.
func TotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err == nil {
		return in.Totalram
	}
	return 0
}
