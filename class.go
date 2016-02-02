package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type Class C.Class

type Imp C.IMP

func Class_getName(cls Class) string {
	cname := C.class_getName(cls)
	return C.GoString(cname)
}

func Class_getSuperclass(cls Class) Class {
	return Class(C.class_getSuperclass(cls))
}

func Class_isMetaClass(cls Class) bool {
	return C.class_isMetaClass(cls) != 0
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
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Ivar(C.class_getClassVariable(cls, cname))
}

func Class_addIvar(cls Class, name string, size uint, alignment uint8, types string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ctypes := C.CString(types)
	defer C.free(unsafe.Pointer(ctypes))

	return C.class_addIvar(cls, cname, C.size_t(size), C.uint8_t(alignment), ctypes) != 0
}

func Class_copyIvarList(cls Class) (ivars []Ivar) {
	var coutCount C.uint

	ivarList := C.class_copyIvarList(cls, &coutCount)
	defer C.free(unsafe.Pointer(ivarList))

	if outCount := uint(coutCount); outCount > 0 {
		ivars = make([]Ivar, outCount)

		for i, elem := uint(0), ivarList; i < outCount; i++ {
			ivars[i] = Ivar(*elem)
			elem = nextIvar(elem)
		}
	}

	return
}

func Class_getProperty(cls Class, name string) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Property(C.class_getProperty(cls, cname))
}

func Class_copyPropertyList(cls Class) (properties []Property) {
	var coutCount C.uint

	propertyList := C.class_copyPropertyList(cls, &coutCount)
	defer C.free(unsafe.Pointer(propertyList))

	if outCount := uint(coutCount); outCount > 0 {
		properties = make([]Property, outCount)

		for i, elem := uint(0), propertyList; i < outCount; i++ {
			properties[i] = Property(*elem)
			elem = nextProperty(elem)
		}
	}

	return
}

func Class_addMethod(cls Class, name Sel, imp Imp, types string) bool {
	ctype := C.CString(types)
	defer C.free(unsafe.Pointer(ctype))

	return C.class_addMethod(cls, name, imp, ctype) != 0
}

func Class_getInstanceMethod(aClass Class, aSelector Sel) Method {
	return Method(C.class_getInstanceMethod(aClass, aSelector))
}

func Class_getClassMethod(aClass Class, aSelector Sel) Method {
	return Method(C.class_getClassMethod(aClass, aSelector))
}

func Class_copyMethodList(cls Class) (methods []Method) {
	var coutCount C.uint

	methodList := C.class_copyMethodList(cls, &coutCount)
	defer C.free(unsafe.Pointer(methodList))

	if outCount := uint(coutCount); outCount > 0 {
		methods = make([]Method, outCount)

		for i, elem := uint(0), methodList; i < outCount; i++ {
			methods[i] = Method(*elem)
			elem = nextMethod(elem)
		}
	}

	return
}

func Class_replaceMethod(cls Class, name Sel, imp Imp, types string) Imp {
	ctype := C.CString(types)
	defer C.free(unsafe.Pointer(ctype))

	return Imp(C.class_replaceMethod(cls, name, imp, ctype))
}

func Class_getMethodImplementation(cls Class, name Sel) Imp {
	return Imp(C.class_getMethodImplementation(cls, name))
}

func Class_getMethodImplementation_stret(cls Class, name Sel) Imp {
	return Imp(C.class_getMethodImplementation_stret(cls, name))
}

func Class_respondsToSelector(cls Class, sel Sel) bool {
	return C.class_respondsToSelector(cls, sel) != 0
}

func Class_addProtocol(cls Class, protocol Protocol) bool {
	return C.class_addProtocol(cls, protocol) != 0
}

func Class_addProperty(cls Class, name string, attributes []PropertyAttribute) bool {
	var cattributes *C.objc_property_attribute_t

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	attrSize := unsafe.Sizeof(*cattributes)
	attributeCount := len(attributes)

	if len(attributes) != 0 {
		cattributes = (*C.objc_property_attribute_t)(calloc(uint(attributeCount), attrSize))

		defer func(cattributes *C.objc_property_attribute_t, attributeCount int) {
			for i, elem := 0, cattributes; i < attributeCount; i++ {
				C.free(unsafe.Pointer(elem.name))
				C.free(unsafe.Pointer(elem.value))

				elem = nextPropertyAttr(elem)
			}

			C.free(unsafe.Pointer(cattributes))
		}(cattributes, attributeCount)

		for i, elem := 0, cattributes; i < attributeCount; i++ {
			attr := attributes[i]
			elem.name = C.CString(attr.Name)
			elem.value = C.CString(attr.Value)
			elem = nextPropertyAttr(elem)
		}
	}

	return C.class_addProperty(cls, cname, cattributes, C.uint(attributeCount)) != 0
}

func Class_replaceProperty(cls Class, name string, attributes []PropertyAttribute) {
	var cattributes *C.objc_property_attribute_t

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	attrSize := unsafe.Sizeof(*cattributes)
	attributeCount := len(attributes)

	if len(attributes) != 0 {
		cattributes = (*C.objc_property_attribute_t)(calloc(uint(attributeCount), attrSize))

		defer func(cattributes *C.objc_property_attribute_t, attributeCount int) {
			for i, elem := 0, cattributes; i < attributeCount; i++ {
				C.free(unsafe.Pointer(elem.name))
				C.free(unsafe.Pointer(elem.value))

				elem = nextPropertyAttr(elem)
			}

			C.free(unsafe.Pointer(cattributes))
		}(cattributes, attributeCount)

		for i, elem := 0, cattributes; i < attributeCount; i++ {
			attr := attributes[i]
			elem.name = C.CString(attr.Name)
			elem.value = C.CString(attr.Value)
			elem = nextPropertyAttr(elem)
		}
	}

	C.class_replaceProperty(cls, cname, cattributes, C.uint(attributeCount))
}

func Class_conformsToProtocol(cls Class, protocol Protocol) bool {
	return C.class_conformsToProtocol(cls, protocol) != 0
}

func Class_copyProtocolList(cls Class) (protocols []Protocol) {
	var coutCount C.uint

	protocolList := C.class_copyProtocolList(cls, &coutCount)
	C.free(unsafe.Pointer(protocolList))

	if outCount := uint(coutCount); outCount > 0 {
		protocols = make([]Protocol, outCount)

		for i, elem := uint(0), protocolList; i < outCount; i++ {
			protocols[i] = Protocol(*elem)
			elem = nextProtocol(elem)
		}
	}

	return
}

func Class_getVersion(theClass Class) int {
	return int(C.class_getVersion(theClass))
}

func Class_setVersion(theClass Class, version int) {
	C.class_setVersion(theClass, C.int(version))
}

func Class_createInstance(cls Class, extraBytes uint) Id {
	return Id(C.class_createInstance(cls, C.size_t(extraBytes)))
}

func Class_getImageName(cls Class) string {
	return C.GoString(C.class_getImageName(cls))
}

func nextClass(list *C.Class) *C.Class {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.Class)(unsafe.Pointer(ptr))
}
