package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

func CBool(value bool) C.BOOL {
	if value {
		return 1
	}

	return 0
}

func calloc(count uint, size uintptr) unsafe.Pointer {
	return C.calloc(C.size_t(count), C.size_t(size))
}

func free(ptr unsafe.Pointer) {
	C.free(ptr)
}
