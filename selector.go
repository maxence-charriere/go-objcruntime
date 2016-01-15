package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type Sel C.SEL

func Sel_getName(aSelector Sel) string {
	return C.GoString(C.sel_getName(aSelector))
}

func Sel_registerName(str string) Sel {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Sel(C.sel_registerName(cstr))
}

func Sel_getUid(str string) Sel {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Sel(C.sel_getUid(cstr))
}

func Sel_isEqual(lhs Sel, rhs Sel) bool {
	return C.sel_isEqual(lhs, rhs) != 0
}
