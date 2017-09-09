// Package memory provides a single method reporting total system memory.
package memory

// TotalMemory returns the total system memory in bytes, or 0 if
// available memory could not be determined.
func TotalMemory() uint64
