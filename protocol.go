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

func Protocol_addMethodDescription(proto Protocol, name Sel, types string, isRequiredMethod bool, isInstanceMethod bool) {
	ctypes := C.CString(types)
	defer C.free(unsafe.Pointer(ctypes))

	C.protocol_addMethodDescription(proto, name, ctypes, cBool(isRequiredMethod), cBool(isInstanceMethod))
}

func Protocol_addProtocol(proto Protocol, addition Protocol) {
	C.protocol_addProtocol(proto, addition)
}

func Protocol_addProperty(proto Protocol, name string, attributes []PropertyAttribute, isRequiredProperty bool, isInstanceProperty bool) {
	var attrPtr *C.objc_property_attribute_t

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	attrSize := unsafe.Sizeof(*attrPtr)
	attributeCount := len(attributes)

	cattributes := (*C.objc_property_attribute_t)(C.malloc(C.size_t(attrSize) * C.size_t(attributeCount)))
	defer C.free(unsafe.Pointer(cattributes))

	elem := cattributes

	for i := 0; i < attributeCount; i++ {
		attr := attributes[i]
		elem.name = C.CString(attr.Name)
		elem.value = C.CString(attr.Value)

		defer C.free(unsafe.Pointer(elem.name))
		defer C.free(unsafe.Pointer(elem.value))

		elem = nextPropertyAttr(elem)
	}

	C.protocol_addProperty(proto, cname, cattributes, C.uint(attributeCount), cBool(isRequiredProperty), cBool(isInstanceProperty))
}

func Protocol_getName(p Protocol) string {
	return C.GoString(C.protocol_getName(p))
}

func Protocol_isEqual(proto Protocol, other Protocol) bool {
	return C.protocol_isEqual(proto, other) != 0
}

func Protocol_copyMethodDescriptionList(p Protocol, isRequiredMethod bool, isInstanceMethod bool) (descriptions []MethodDescription) {
	var coutCount C.uint

	descriptionList := C.protocol_copyMethodDescriptionList(p, cBool(isRequiredMethod), cBool(isInstanceMethod), &coutCount)
	defer C.free(unsafe.Pointer(descriptionList))

	if outCount := uint(coutCount); outCount > 0 {
		descriptions = make([]MethodDescription, outCount)
		elem := descriptionList

		for i := uint(0); i < outCount; i++ {
			descriptions[i] = makeMethodDescription(*elem)

			defer C.free(unsafe.Pointer(elem.types))

			elem = nextMethodDescription(elem)
		}
	}

	return
}

func Protocol_getMethodDescription(p Protocol, aSel Sel, isRequiredMethod bool, isInstanceMethod bool) MethodDescription {
	cmethodDescription := C.protocol_getMethodDescription(p, aSel, cBool(isRequiredMethod), cBool(isInstanceMethod))
	return makeMethodDescription(cmethodDescription)
}

func Protocol_copyPropertyList(protocol Protocol) (properties []Property) {
	var coutCount C.uint

	propertyList := C.protocol_copyPropertyList(protocol, &coutCount)
	defer C.free(unsafe.Pointer(propertyList))

	if outCount := uint(coutCount); outCount > 0 {
		properties = make([]Property, outCount)
		elem := propertyList

		for i := uint(0); i < outCount; i++ {
			properties[i] = Property(*elem)
			elem = nextProperty(elem)
		}
	}

	return
}

func Protocol_getProperty(proto Protocol, name string, isRequiredProperty bool, isInstanceProperty bool) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Property(C.protocol_getProperty(proto, cname, cBool(isRequiredProperty), cBool(isInstanceProperty)))
}

func Protocol_copyProtocolList(proto Protocol) (protocols []Protocol) {
	var coutCount C.uint

	protocolList := C.protocol_copyProtocolList(proto, &coutCount)
	defer C.free(unsafe.Pointer(protocolList))

	if outCount := uint(coutCount); outCount > 0 {
		protocols = make([]Protocol, outCount)
		elem := protocolList

		for i := uint(0); i < outCount; i++ {
			protocols[i] = Protocol(*elem)
			elem = nextProtocol(elem)
		}
	}

	return
}

func Protocol_conformsToProtocol(proto Protocol, other Protocol) bool {
	return C.protocol_conformsToProtocol(proto, other) != 0
}

func nextProtocol(list **C.Protocol) **C.Protocol {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (**C.Protocol)(unsafe.Pointer(ptr))
}
