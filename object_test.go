package objc

import (
	"testing"
	"unsafe"
)

func TestObjectCopy(t *testing.T) {
	class := Objc_getClass("NSObject")
	objA := Class_createInstance(class, 0)

	if objB := Object_copy(objA, 0); objA == objB {
		t.Errorf("objects should be different: %p == %p", objA, objB)
	}
}

func TestObjectDispose(t *testing.T) {
	class := Objc_getClass("NSObject")
	obj := Class_createInstance(class, 0)

	if obj = Object_dispose(obj); obj != nil {
		t.Errorf("obj should be nil: %#v", obj)
	}
}

func TestObjectSetInstanceVariable(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassSetInstanceVariable"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)
	ivarSize := uint(unsafe.Sizeof(ivarValue))
	Class_addIvar(class, ivarName, ivarSize, 0, "i")

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)

	if ivar := Object_setInstanceVariable(instance, ivarName, unsafe.Pointer(&ivarValue)); ivar == nil {
		t.Error("ivar should not be nil")
	}
}

func TestObjectSetNonexistentInstanceVariable(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassSetNonexistentIvar"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)

	if ivar := Object_setInstanceVariable(instance, ivarName, unsafe.Pointer(&ivarValue)); ivar != nil {
		t.Errorf("ivar should be nil: %#v", ivar)
	}
}

func TestObjectGetInstanceVariable(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassGetInstanceVariable"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)
	ivarSize := uint(unsafe.Sizeof(ivarValue))
	Class_addIvar(class, ivarName, ivarSize, 0, "i")

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)
	setIvar := Object_setInstanceVariable(instance, ivarName, unsafe.Pointer(&ivarValue))

	ivar, outValue := Object_getInstanceVariable(instance, ivarName)

	if ivar != setIvar {
		t.Errorf("ivar should be equal to setIvar: %#v != %#v", ivar, setIvar)
	}

	if value := *(*int32)(outValue); value != ivarValue {
		t.Errorf("value should be equal %d: %d", ivarValue, value)
	}
}

func TestObjectGetNonexistentInstanceVariable(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassGetNonexistentInstanceVariable"
	class := Objc_allocateClassPair(nsobject, className, 0)
	ivarName := "ivar"

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)

	ivar, outValue := Object_getInstanceVariable(instance, ivarName)

	if ivar != nil {
		t.Error("ivar should be nil")
	}

	if outValue != nil {
		t.Errorf("outValue should be nil")
	}
}

func TestObjectGetIndexedIvars(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassGetIndexedIvar"
	class := Objc_allocateClassPair(nsobject, className, 0)
	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 42)

	if index := Object_getIndexedIvars(instance); index == nil {
		t.Error("index should not be nil")
	}
}

func TestObjectGetIvar(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassGetIvar"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)
	ivarSize := uint(unsafe.Sizeof(ivarValue))
	Class_addIvar(class, ivarName, ivarSize, 0, "i")

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)
	ivar := Object_setInstanceVariable(instance, ivarName, unsafe.Pointer(&ivarValue))

	rawValue := Object_getIvar(instance, ivar)

	if value := *(*int32)(rawValue); value != ivarValue {
		t.Errorf("value should be %d: %d", ivarValue, value)
	}
}

func TestObjectGetNonexistentIvar(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassGetNonexistentIvar"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)
	ivarSize := uint(unsafe.Sizeof(ivarValue))
	Class_addIvar(class, ivarName, ivarSize, 0, "i")
	ivar := Class_getInstanceVariable(class, ivarName)

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)

	if value := Object_getIvar(instance, ivar); value != nil {
		t.Error("value should be nil")
	}
}

func TestObjectGetNonexistentInstanceIvar(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassGetNonexistentInstanceIvar"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)
	ivarSize := uint(unsafe.Sizeof(ivarValue))
	Class_addIvar(class, ivarName, ivarSize, 0, "i")
	ivar := Class_getInstanceVariable(class, ivarName)

	Objc_registerClassPair(class)

	if value := Object_getIvar(nil, ivar); value != nil {
		t.Error("value should be nil")
	}
}

func TestObjectsetIvar(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassSetIvar"
	class := Objc_allocateClassPair(nsobject, className, 0)

	ivarName := "ivar"
	ivarValue := int32(42)
	ivarSize := uint(unsafe.Sizeof(ivarValue))
	Class_addIvar(class, ivarName, ivarSize, 0, "i")

	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)
	ivar := Class_getInstanceVariable(class, ivarName)
	Object_setIvar(instance, ivar, unsafe.Pointer(&ivarValue))

	rawValue := Object_getIvar(instance, ivar)

	if value := *(*int32)(rawValue); value != ivarValue {
		t.Errorf("value should be %d: %d", ivarValue, value)
	}
}

func TestObjectGetClassName(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassNameToGetByInstance"
	class := Objc_allocateClassPair(nsobject, className, 0)
	Objc_registerClassPair(class)
	instance := Class_createInstance(class, 0)

	if name := Object_getClassName(instance); name != className {
		t.Errorf("name should be %s: %s", className, name)
	}
}

func TestObjectGetClass(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "ClassToGetByInstance"
	class := Objc_allocateClassPair(nsobject, className, 0)
	Objc_registerClassPair(class)
	instance := Class_createInstance(class, 0)

	if getClass := Object_getClass(instance); class != getClass {
		t.Errorf("getClass should be equal to class: %#v != %#v", getClass, class)
	}
}

func TestObjectSetClass(t *testing.T) {
	nsobject := Objc_getClass("NSObject")
	className := "InstanceToSetClass"
	class := Objc_allocateClassPair(nsobject, className, 0)
	Objc_registerClassPair(class)
	instance := Class_createInstance(class, 0)

	if oldClass := Object_setClass(instance, nsobject); oldClass != class {
		t.Errorf("oldClass should be equal to class: %#v != %#v", oldClass, class)
	}
}

func TestObjectSetNilInstanceClass(t *testing.T) {
	nsobject := Objc_getClass("NSObject")

	if oldClass := Object_setClass(nil, nsobject); oldClass != nil {
		t.Error("oldClass should be nil")
	}
}
