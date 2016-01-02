package objcruntime

// #include <stdlib.h>
// #include <objc/runtime.h>
//
// static Ivar *ivar_offset(Ivar *p, size_t n) {
//   return p + n;
// }
import "C"
import "unsafe"

// Class define an opaque type that represents an Objective-C class.
type Class C.Class

// ClassGetName the name of a class.
func ClassGetName(cls Class) string {
	cname := C.class_getName(cls)
	return C.GoString(cname)
}

// ClassGetSuperclass returns the superclass of a class.
//
// Discussion:
// - You should usually use NSObject‘s superclass method instead of this function.
func ClassGetSuperclass() (cls Class) {
	return Class(C.class_getSuperclass(cls))
}

// ClassIsMetaClass returns a Boolean value that indicates whether a class object is a metaclass.
func ClassIsMetaClass(cls Class) bool {
	return C.class_isMetaClass(cls) == 1
}

// ClassGetInstanceSize returns the size of instances of a class.
func ClassGetInstanceSize(cls Class) uint {
	return uint(C.class_getInstanceSize(cls))
}

// ClassGetInstanceVariable returns the Ivar for a specified instance variable of a given class.
func ClassGetInstanceVariable(cls Class, name string) Ivar {
	cname := C.CString(string(name))
	defer C.free(unsafe.Pointer(cname))

	return Ivar(C.class_getInstanceVariable(cls, cname))
}

// ClassGetClassVariable returns the Ivar for a specified class variable of a given class.
func ClassGetClassVariable(cls Class, name string) Ivar {
	cname := C.CString(string(name))
	defer C.free(unsafe.Pointer(cname))

	return Ivar(C.class_getClassVariable(cls, cname))
}

// ClassAddIvar adds a new instance variable to a class.
//
// Discussion:
// - This function may only be called after ObjcAllocateClassPair and before ObjcRegisterClassPair.
//   Adding an instance variable to an existing class is not supported.
// - The class must not be a metaclass.
//   Adding an instance variable to a metaclass is not supported.
// - The instance variable's minimum alignment in bytes is 1<<align.
//   The minimum alignment of an instance variable depends on the ivar's type and the machine architecture.
//   For variables of any pointer type, pass log2(sizeof(pointer_type)).
func ClassAddIvar(cls Class, name string, size uint, alignment uint8, types string) bool {
	cname := C.CString(string(name))
	defer C.free(unsafe.Pointer(cname))

	ctypes := C.CString(string(types))
	defer C.free(unsafe.Pointer(ctypes))

	return C.class_addIvar(cls, cname, C.size_t(size), C.uint8_t(alignment), ctypes) == 1
}

// ClassCopyIvarList describes the instance variables declared by a class.
func ClassCopyIvarList(cls Class) (ivars []Ivar, outCount uint) {
	var coutCount C.uint

	ivarListPtr := C.class_copyIvarList(cls, &coutCount)
	defer C.free(unsafe.Pointer(ivarListPtr))

	outCount = uint(coutCount)

	if outCount > 0 {
		ivars = make([]Ivar, 0, outCount)

		for i := uint(0); i < outCount; i++ {
			ivarOffset := C.ivar_offset(ivarListPtr, C.size_t(i))
			ivars = append(ivars, Ivar(*ivarOffset))
		}
	}

	return
}
