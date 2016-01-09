package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type Property C.objc_property_t

type PropertyAttribute struct {
	Name  string
	Value string
}

func makePropertyAttribute(attr C.objc_property_attribute_t) PropertyAttribute {
	return PropertyAttribute{
		Name:  C.GoString(attr.name),
		Value: C.GoString(attr.value),
	}
}

func Propety_getName(property Property) string {
	cname := C.property_getName(property)
	return C.GoString(cname)
}

func Property_getAttributes(property Property) string {
	cattr := C.property_getAttributes(property)
	return C.GoString(cattr)
}

func Property_copyAttributeValue(property Property, attributeName string) string {
	cattrName := C.CString(attributeName)
	defer C.free(unsafe.Pointer(cattrName))

	cattrVal := C.property_copyAttributeValue(property, cattrName)
	defer C.free(unsafe.Pointer(cattrVal))

	return C.GoString(cattrVal)
}

func Property_copyAttributeList(property Property) (attributes []PropertyAttribute) {
	var coutCount C.uint

	attrList := C.property_copyAttributeList(property, &coutCount)
	defer C.free(unsafe.Pointer(attrList))

	if outCount := uint(coutCount); outCount > 0 {
		attributes = make([]PropertyAttribute, outCount)
		elem := attrList

		for i := uint(0); i < outCount; i++ {
			attributes[i] = makePropertyAttribute(*elem)

			defer C.free(unsafe.Pointer(elem.name))
			defer C.free(unsafe.Pointer(elem.value))

			elem = nextPropertyAttr(elem)
		}
	}

	return
}

func nextProperty(list *C.objc_property_t) *C.objc_property_t {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.objc_property_t)(unsafe.Pointer(ptr))
}

func nextPropertyAttr(list *C.objc_property_attribute_t) *C.objc_property_attribute_t {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.objc_property_attribute_t)(unsafe.Pointer(ptr))
}
