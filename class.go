package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
//
// static Ivar *ivar_offset(Ivar *p, size_t n) {
//   return p + n;
// }
import "C"
import "unsafe"

type Class C.Class

func Class_getName(cls Class) string {
	cname := C.class_getName(cls)
	return C.GoString(cname)
}

func Class_getSuperclass(cls Class) Class {
	return Class(C.class_getSuperclass(cls))
}

func Class_isMetaClass(cls Class) bool {
	return C.class_isMetaClass(cls) == 1
}

func Class_getInstanceSize(cls Class) uint {
	return uint(C.class_getInstanceSize(cls))
}

func Class_getInstanceVariable(cls Class, name string) Ivar {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Ivar(C.class_getInstanceVariable(cls, cname))
}

func Class_getClassVariable(cls Class, name string) Ivar {
	cname := C.CString(string(name))
	defer C.free(unsafe.Pointer(cname))

	return Ivar(C.class_getClassVariable(cls, cname))
}

func Class_addIvar(cls Class, name string, size uint, alignment uint8, types string) bool {
	cname := C.CString(string(name))
	defer C.free(unsafe.Pointer(cname))

	ctypes := C.CString(string(types))
	defer C.free(unsafe.Pointer(ctypes))

	return C.class_addIvar(cls, cname, C.size_t(size), C.uint8_t(alignment), ctypes) == 1
}

func Class_copyIvarList(cls Class) (ivarList []Ivar) {
	var coutCount C.uint

	ivarListPtr := C.class_copyIvarList(cls, &coutCount)
	defer C.free(unsafe.Pointer(ivarListPtr))

	if outCount := uint(coutCount); outCount > 0 {
		ivarList = make([]Ivar, outCount)

		for i := uint(0); i != outCount; i++ {
			ivarOffset := C.ivar_offset(ivarListPtr, C.size_t(i))
			ivarList[i] = Ivar(*ivarOffset)
		}
	}

	return
}
