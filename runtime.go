package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
//
// static Method *method_offset(Method *p, size_t n) {
//   return p + n;
// }
import "C"
import "unsafe"

type Class C.Class

type Method C.Method

type Ivar C.Ivar

type Property C.objc_property_t

type Sel C.SEL

type Id C.id

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

func Class_copyIvarList(cls Class) (ivarList []Ivar) {
	var coutCount C.uint

	list := C.class_copyIvarList(cls, &coutCount)
	defer C.free(unsafe.Pointer(list))

	if outCount := uint(coutCount); outCount > 0 {
		ivarList = make([]Ivar, outCount)
		elem := list

		for i := uint(0); i < outCount; i++ {
			ivarList[i] = Ivar(*elem)
			elem = nextIvar(elem)
		}
	}

	return
}

func Class_getIvarLayout(cls Class) uintptr {
	return uintptr(unsafe.Pointer(C.class_getIvarLayout(cls)))
}

func Class_setIvarLayout(cls Class, layout uintptr) {
	C.class_setIvarLayout(cls, (*C.uint8_t)(unsafe.Pointer(layout)))
}

func Class_getWeakIvarLayout(cls Class) uintptr {
	return uintptr(unsafe.Pointer(C.class_getWeakIvarLayout(cls)))
}

func Class_setWeakIvarLayout(cls Class, layout uintptr) {
	C.class_setWeakIvarLayout(cls, (*C.uint8_t)(unsafe.Pointer(layout)))
}

func Class_getProperty(cls Class, name string) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return Property(C.class_getProperty(cls, cname))
}

func Class_copyPropertyList(cls Class) (properties []Property) {
	var coutCount C.uint

	list := C.class_copyPropertyList(cls, &coutCount)
	defer C.free(unsafe.Pointer(list))

	if outCount := uint(coutCount); outCount > 0 {
		properties := make([]Property, outCount)
		elem := list

		for i := uint(0); i < outCount; i++ {
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

func Class_getClassMethodt(aClass Class, aSelector Sel) Method {
	return Method(C.class_getClassMethod(aClass, aSelector))
}

func Class_copyMethodList(cls Class) (methods []Method) {
	var coutCount C.uint

	list := C.class_copyMethodList(cls, &coutCount)
	defer C.free(unsafe.Pointer(list))

	if outCount := uint(coutCount); outCount > 0 {
		methods := make([]Method, outCount)
		elem := list

		for i := uint(0); i < outCount; i++ {
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

// SEL
func Sel_getName(aSelector Sel) string {
	return C.GoString(C.sel_getName(aSelector))
}

func Sel_registerName(str string) Sel {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Sel(C.sel_registerName(cstr))
}

func Sel_getUid(str string) Sel {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return Sel(C.sel_getUid(cstr))
}

func Sel_isEqual(lhs Sel, rhs Sel) bool {
	return C.sel_isEqual(lhs, rhs) != 0
}

// Helpers
func nextIvar(list *C.Ivar) *C.Ivar {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.Ivar)(unsafe.Pointer(ptr))
}

func nextProperty(list *C.objc_property_t) *C.objc_property_t {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.objc_property_t)(unsafe.Pointer(ptr))
}

func nextMethod(list *C.Method) *C.Method {
	ptr := uintptr(unsafe.Pointer(list)) + unsafe.Sizeof(*list)
	return (*C.Method)(unsafe.Pointer(ptr))
}
