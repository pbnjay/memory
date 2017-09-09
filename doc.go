// +build !linux,!darwin,!windows

// Package memory provides a single method reporting total system memory.
package memory

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	return 0
}
