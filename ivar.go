package objc

// #include <objc/runtime.h>
import "C"
import "unsafe"

type Ivar C.Ivar

func Ivar_getName(ivar Ivar) string {
	return C.GoString(C.ivar_getName(ivar))
}

func Ivar_getTypeEncoding(ivar Ivar) string {
	return C.GoString(C.ivar_getTypeEncoding(ivar))
}

func nextIvar(list *C.Ivar) *C.Ivar {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.Ivar)(unsafe.Pointer(ptr))
}
