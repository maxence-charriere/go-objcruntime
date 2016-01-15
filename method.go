package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type Method C.Method

type MethodDescription struct {
	Name  Sel
	Types string
}

func makeMethodDescription(description C.struct_objc_method_description) MethodDescription {
	return MethodDescription{
		Name:  Sel(description.name),
		Types: C.GoString(description.types),
	}
}

func nextMethod(list *C.Method) *C.Method {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.Method)(unsafe.Pointer(ptr))
}

func nextMethodDescription(list *C.struct_objc_method_description) *C.struct_objc_method_description {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.struct_objc_method_description)(unsafe.Pointer(ptr))
}
