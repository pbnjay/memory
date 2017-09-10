// +build darwin

package memory

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	s, err := sysctlUint64("hw.memsize")
	if err != nil {
		return 0
	}
	return s
}
