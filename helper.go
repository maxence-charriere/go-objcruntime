package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

func cBool(value bool) C.BOOL {
	if value {
		return 1
	}

	return 0
}

func calloc(count uint, size uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.calloc(C.size_t(count), C.size_t(size)))
}
