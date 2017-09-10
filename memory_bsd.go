// +build netbsd openbsd
// +build amd64

package memory

import (
	"encoding/binary"
	"syscall"
)

// TotalMemory returns the total system memory in bytes, or 0 if
// installed memory size could not be determined.
func TotalMemory() uint64 {
	// FIXME: there is no 64bit version like this:
	//s, err := syscall.SysctlUint64("hw.physmem64")

	s, err := syscall.Sysctl("hw.physmem64")
	if err != nil {
		return 0
	}
	// hack because the string conversion above drops \0
	b := []byte(s)
	for len(b) < 8 {
		b = append(b, 0)
	}
	return binary.LittleEndian.Uint64(b)
}
