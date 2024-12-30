//go:build windows
// +build windows

package local

import (
	"syscall"
	"unsafe"

	"github.com/alist-org/alist/v3/internal/model"
)

func getDiskSpace(path string) model.Space {
	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64

	// Load the DLL
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getDiskFreeSpaceExW := kernel32.NewProc("GetDiskFreeSpaceExW")

	// Convert path to UTF-16
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return model.Space{}
	}

	// Call the function
	ret, _, err := getDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
	)

	// Check the result
	if ret == 0 {
		return model.Space{}
	}

	return model.Space{
		Usage: int64(totalNumberOfBytes - totalNumberOfFreeBytes),
		Total: int64(totalNumberOfBytes),
	}
}
