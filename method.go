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

func Method_getName(method Method) Sel {
	return Sel(C.method_getName(method))
}

func Method_getImplementation(method Method) Imp {
	return Imp(C.method_getImplementation(method))
}

func Method_getTypeEncoding(method Method) string {
	return C.GoString(C.method_getTypeEncoding(method))
}

func Method_copyReturnType(method Method) string {
	return C.GoString(C.method_copyReturnType(method))
}

func Method_copyArgumentType(method Method, index uint) string {
	return C.GoString(C.method_copyArgumentType(method, C.uint(index)))
}

func Method_getNumberOfArguments(method Method) uint {
	return uint(C.method_getNumberOfArguments(method))
}

func Method_getDescription(m Method) MethodDescription {
	return makeMethodDescription(*C.method_getDescription(m))
}

func Method_setImplementation(method Method, imp Imp) Imp {
	return Imp(C.method_setImplementation(method, imp))
}

func Method_exchangeImplementations(m1 Method, m2 Method) {
	C.method_exchangeImplementations(m1, m2)
}

func nextMethod(list *C.Method) *C.Method {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.Method)(unsafe.Pointer(ptr))
}

func nextMethodDescription(list *C.struct_objc_method_description) *C.struct_objc_method_description {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.struct_objc_method_description)(unsafe.Pointer(ptr))
}
