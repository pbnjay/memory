package memory

// TotalMemory returns the total system memory in bytes, or 0 if
// available memory could not be determined.
func TotalMemory() uint64 {

	/*
	   #include <Windows.h>
	   ....
	   uint64_t memory;
	   GetPhysicallyInstalledSystemMemory(&memory);


	*/

	return 0
}
