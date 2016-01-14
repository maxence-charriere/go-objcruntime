package objc

import (
	"testing"

	"github.com/achille-roussel/go-ffi"
)

func TestClassGetName(t *testing.T) {
	name := "NSObject"
	class := Objc_getClass(name)

	if n := Class_getName(class); n != name {
		t.Errorf("class should be named %s: %s", name, n)
	}
}

func TestClassGetSuperclass(t *testing.T) {
	var superclass Class

	nsObject := Objc_getClass("NSObject")
	class := Objc_allocateClassPair(nsObject, "GetClassSuperClass", 0)

	if superclass = Class_getSuperclass(class); superclass == nil {
		t.Fatal("superclass should not be nil")
	}

	if superclass != nsObject {
		t.Errorf("super class should be an ns object: %#v != %#v", superclass, nsObject)
	}
}

func TestClassGetNonexistentSuperclass(t *testing.T) {
	nsObject := Objc_getClass("NSObject")

	if superclass := Class_getSuperclass(nsObject); superclass != nil {
		t.Fatalf("superclass should be nil: %#v", superclass)
	}
}

func TestClassIsMetaClass(t *testing.T) {
	nsObject := Objc_getMetaClass("NSObject")

	if !Class_isMetaClass(nsObject) {
		t.Errorf("nsObject should be a meta class: %#v", nsObject)
	}
}

func TestClassIsNotMetaClass(t *testing.T) {
	nsObject := Objc_getClass("NSObject")

	if Class_isMetaClass(nsObject) {
		t.Errorf("nsObject should not be a meta class: %#v", nsObject)
	}
}

func TestClassGetInstanceSize(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	if size := Class_getInstanceSize(nsObject); size == 0 {
		t.Error("size should not be 0")
	}
}

func TestClassGetCustomInstanceSize(t *testing.T) {
	className := "CustomClassForSize"
	class := Objc_allocateClassPair(nil, className, 0)

	Class_addIvar(class, "a", 4, 0, "i")
	Class_addIvar(class, "b", 4, 0, "i")
	Class_addIvar(class, "c", 4, 0, "i")

	if size := Class_getInstanceSize(class); size != 24 {
		t.Errorf("size should be equal to 24: %d", size)
	}
}

func TestClassGetInstanceVariable(t *testing.T) {
	className := "GetIvarClass"
	class := Objc_allocateClassPair(nil, className, 0)
	ivarName := "ivar"
	Class_addIvar(class, ivarName, 4, 4, "i")

	if ivar := Class_getInstanceVariable(class, ivarName); ivar == nil {
		t.Error("ivar should not be nil")
	}
}

func TestClassGetClassVariable(t *testing.T) {
	className := "GetCvarClass"
	class := Objc_allocateClassPair(nil, className, 0)
	ivarName := "ivar"
	Class_addIvar(class, ivarName, 4, 0, "i")

	if ivar := Class_getClassVariable(class, ivarName); ivar == nil {
		t.Error("ivar should not be nil")
	}
}

func TestClassAddIvar(t *testing.T) {
	className := "AddIvarClass"
	class := Objc_allocateClassPair(nil, className, 0)
	ivarName := "ivar"

	if !Class_addIvar(class, ivarName, 4, 0, "i") {
		t.Errorf("add %s to %s failed", ivarName, className)
	}
}

func TestClassAddSameIvar(t *testing.T) {
	className := "AddIvarClass"
	class := Objc_allocateClassPair(nil, className, 0)
	ivarName := "ivar"
	Class_addIvar(class, ivarName, 4, 0, "i")

	if Class_addIvar(class, ivarName, 4, 0, "i") {
		t.Errorf("add %s to %s should have failed", ivarName, className)
	}
}

func TestClassCopyIvarList(t *testing.T) {
	className := "ClassForCopyIvar"
	class := Objc_allocateClassPair(nil, className, 0)

	Class_addIvar(class, "a", 4, 0, "i")
	Class_addIvar(class, "b", 4, 0, "i")
	Class_addIvar(class, "c", 4, 0, "i")

	ivars := Class_copyIvarList(class)

	if l := len(ivars); l != 3 {
		t.Errorf("ivars should have a len of 3: %d", l)
	}
}

func TestClassCopyEmptyIvarList(t *testing.T) {
	className := "ClassWithoutIvars"
	class := Objc_allocateClassPair(nil, className, 0)

	ivars := Class_copyIvarList(class)

	if l := len(ivars); l != 0 {
		t.Errorf("ivars should be empty:", l)
	}
}

func TestClassGetProperty(t *testing.T) {
	className := "ClassForGetProperty"
	class := Objc_allocateClassPair(nil, className, 0)
	propertyName := "A"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Class_addProperty(class, propertyName, attributes)

	if property := Class_getProperty(class, propertyName); property == nil {
		t.Error("property should not be nil")
	}
}

func TestClassGetNonexistentProperty(t *testing.T) {
	nsObject := Objc_getClass("NSObject")

	if property := Class_getProperty(nsObject, "A"); property != nil {
		t.Errorf("property should be nil: %#v", property)
	}
}

func TestClassCopyPropertyList(t *testing.T) {
	className := "ClassForCopyPropertyList"
	class := Objc_allocateClassPair(nil, className, 0)

	aAttributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	bAttributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "i"},
	}

	Class_addProperty(class, "A", aAttributes)
	Class_addProperty(class, "B", bAttributes)

	properties := Class_copyPropertyList(class)

	if l := len(properties); l != 2 {
		t.Errorf("properties len should be 2: %d", l)
	}
}

func TestClassCopyEmptyPropertyList(t *testing.T) {
	className := "ClassForCopyEmptyPropertyList"
	class := Objc_allocateClassPair(nil, className, 0)

	properties := Class_copyPropertyList(class)

	if l := len(properties); l != 0 {
		t.Errorf("properties len should be 0: %d", l)
	}
}

func TestClassAddMethod(t *testing.T) {
	//className := "ClassWithMethod"
	//class := Objc_allocateClassPair(nil, className, 0)
	//methodeName := "MethodA"
	method := func(id Id, sel Sel) {}
	closure := ffi.Closure(method)
	methodPtr := closure.Pointer()
	t.Log(methodPtr)
	// sel := Sel_registerName(methodeName)

	// if !Class_addMethod(class, sel, Imp(unsafe.Pointer(methodPtr)), "v") {
	// 	t.Errorf("failed to add %s to %s", methodeName, className)
	// }
}
