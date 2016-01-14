package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

func Objc_allocateClassPair(superclass Class, name string, extraBytes uint) Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Class(C.objc_allocateClassPair(superclass, cname, C.size_t(extraBytes)))
}

func Objc_disposeClassPair(cls Class) {
	C.objc_disposeClassPair(cls)
}

func Objc_registerClassPair(cls Class) {
	C.objc_registerClassPair(cls)
}

func Objc_copyClassList() (classes []Class) {
	var coutCount C.uint

	classList := C.objc_copyClassList(&coutCount)
	defer C.free(unsafe.Pointer(classList))

	if outCount := uint(coutCount); outCount > 0 {
		classes = make([]Class, outCount)
		elem := classList

		for i := uint(0); i < outCount; i++ {
			classes[i] = Class(*elem)
			elem = nextClass(elem)
		}
	}

	return
}

func Objc_getClass(name string) Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Class(C.objc_getClass(cname))
}

func Objc_getMetaClass(name string) Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Class(C.objc_getMetaClass(cname))
}

func Objc_getProtocol(name string) Protocol {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Protocol(C.objc_getProtocol(cname))
}

func Objc_copyProtocolList() (protocols []Protocol) {
	var coutCount C.uint

	protocolList := C.objc_copyProtocolList(&coutCount)
	defer C.free(unsafe.Pointer(protocolList))

	if outCount := uint(coutCount); outCount > 0 {
		protocols = make([]Protocol, outCount)
		elem := protocolList

		for i := uint(0); i < outCount; i++ {
			protocols[i] = *elem
			elem = nextProtocol(elem)
		}
	}

	return
}

func Objc_allocateProtocol(name string) Protocol {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Protocol(C.objc_allocateProtocol(cname))
}

func Objc_registerProtocol(protocol Protocol) {
	C.objc_registerProtocol(protocol)
}
