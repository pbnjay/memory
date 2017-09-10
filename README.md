# memory (WIP)

Provides a simple cross-platform function 
`func memory.TotalMemory() uint64` that returns the
total bytes of system memory installed in a system:

See https://github.com/golang/go/issues/21816 for some justification.

Tested/working on:
 - macOS 10.12.6 (16G29)
 - Windows 10 1511 (10586.1045)
 - Linux RHEL (3.10.0-327.3.1.el7.x86_64)

If you have access to 32bit or BSD systems please
test and file bugs if necessary.