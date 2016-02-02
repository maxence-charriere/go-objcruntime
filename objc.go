package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type AssociationPolicy uintptr

const (
	OBJC_ASSOCIATION_ASSIGN           AssociationPolicy = 0
	OBJC_ASSOCIATION_RETAIN_NONATOMIC AssociationPolicy = 1
	OBJC_ASSOCIATION_COPY_NONATOMIC   AssociationPolicy = 3
	OBJC_ASSOCIATION_RETAIN           AssociationPolicy = 01401
	OBJC_ASSOCIATION_COPY             AssociationPolicy = 01403
)

func Objc_allocateClassPair(superclass Class, name string, extraBytes uint) Class {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	return Class(C.objc_allocateClassPair(superclass, cname, C.size_t(extraBytes)))
}

func Objc_disposeClassPair(cls Class) {
	C.objc_disposeClassPair(cls)
}

func Objc_registerClassPair(cls Class) {
	C.objc_registerClassPair(cls)
}

func Objc_constructInstance(cls Class, bytes unsafe.Pointer) Id {
	return Id(C.objc_constructInstance(cls, bytes))
}

func Objc_destructInstance(obj Id) {
	C.objc_destructInstance(obj)
}

func Objc_copyClassList() (classes []Class) {
	var coutCount C.uint

	classList := C.objc_copyClassList(&coutCount)
	defer free(unsafe.Pointer(classList))

	if outCount := uint(coutCount); outCount > 0 {
		classes = make([]Class, outCount)

		for i, elem := uint(0), classList; i < outCount; i++ {
			classes[i] = Class(*elem)
			elem = nextClass(elem)
		}
	}

	return
}

func Objc_getClass(name string) Class {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	return Class(C.objc_getClass(cname))
}

func Objc_getMetaClass(name string) Class {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	return Class(C.objc_getMetaClass(cname))
}

func Objc_getProtocol(name string) Protocol {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	return Protocol(C.objc_getProtocol(cname))
}

func Objc_copyProtocolList() (protocols []Protocol) {
	var coutCount C.uint

	protocolList := C.objc_copyProtocolList(&coutCount)
	defer free(unsafe.Pointer(protocolList))

	if outCount := uint(coutCount); outCount > 0 {
		protocols = make([]Protocol, outCount)

		for i, elem := uint(0), protocolList; i < outCount; i++ {
			protocols[i] = *elem
			elem = nextProtocol(elem)
		}
	}

	return
}

func Objc_allocateProtocol(name string) Protocol {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	return Protocol(C.objc_allocateProtocol(cname))
}

func Objc_registerProtocol(protocol Protocol) {
	C.objc_registerProtocol(protocol)
}

func Objc_setAssociatedObject(object Id, key unsafe.Pointer, value Id, policy AssociationPolicy) {
	C.objc_setAssociatedObject(object, key, value, C.objc_AssociationPolicy(policy))
}

func Objc_getAssociatedObject(object Id, key unsafe.Pointer) Id {
	return Id(C.objc_getAssociatedObject(object, key))
}

func Objc_removeAssociatedObjects(object Id) {
	C.objc_removeAssociatedObjects(object)
}
