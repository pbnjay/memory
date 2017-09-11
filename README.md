# memory

Package memory provides a single method reporting total system memory
accessible to user code. This value can be used to keep code from thrashing or swapping.

Ideally this would be provided by the stdlib (similar to `runtime.NumCPU`) - Feedback on the proposal is welcome at: https://github.com/golang/go/issues/21816

## `func TotalMemory() uint64`

TotalMemory returns the total accessible system memory in bytes. Note
that accessible memory is total installed physical memory size, minus
reserved areas for the kernel and hardware (if such reservations are
reported by the operating system).

Unlike Available memory, this value is stable during program execution
and does not take into account memory used by other processes.

If accessible memory size could not be determined, then 0 is returned.

## Testing

Tested/working on:
 - macOS 10.12.6 (16G29)
 - Windows 10 1511 (10586.1045)
 - Linux RHEL (3.10.0-327.3.1.el7.x86_64)

Tested on virtual machines:
 - Windows 7 SP1 386
 - Debian stretch 386
 - NetBSD 7.1 amd64 + 386
 - OpenBSD 6.1 amd64 + 386
 - FreeBSD 11.1 amd64 + 386
 - DragonFly BSD 4.8.1 amd64

If you have access to untested systems (notably arm) please
test and file bugs if necessary.