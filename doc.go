// Package memory provides a single method reporting total system memory
// accessible to user code.
package memory

// TotalMemory returns the total accessible system memory in bytes. Note that
// accessible memory is total installed physical memory size, minus reserved
// areas for the kernel and hardware (if such reservations are reported by
// the operating system).
//
// Unlike Available memory, this value is stable during program execution and
// does not take into account memory used by other processes.
//
// If accessible memory size could not be determined, then 0 is returned.
func TotalMemory() uint64 {
	return sysTotalMemory()
}
