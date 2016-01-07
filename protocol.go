package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type Protocol *C.Protocol

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

func nextProtocol(list **C.Protocol) **C.Protocol {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (**C.Protocol)(unsafe.Pointer(ptr))
}
