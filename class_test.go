package objc

import (
	"testing"
	"unsafe"

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
	className := "ClassWithMethod"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)

	if !Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), "v@:") {
		t.Errorf("failed to add %s to %s", methodeName, className)
	}
}

func TestClassAddExistingMethod(t *testing.T) {
	className := "ClassWithMethods"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	closureA := ffi.Closure(func(id Id, sel Sel) {})
	closureB := ffi.Closure(func(id Id, sel Sel, n int) {})
	sel := Sel_registerName(methodeName)

	Class_addMethod(class, sel, Imp(unsafe.Pointer(closureA.Pointer())), "v@:")

	if Class_addMethod(class, sel, Imp(unsafe.Pointer(closureB.Pointer())), "v@:i") {
		t.Errorf("add %s to %s should have failde", methodeName, className)
	}
}

func TestClassGetInstanceMethod(t *testing.T) {
	className := "ClassWithInstanceMethod"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)

	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), "v@:")

	if method := Class_getInstanceMethod(class, sel); method == nil {
		t.Error("method should not be nil")
	}
}

func TestClassGetEmptyInstanceMethod(t *testing.T) {
	className := "ClassWithoutInstanceMethod"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	sel := Sel_registerName(methodeName)

	if method := Class_getInstanceMethod(class, sel); method != nil {
		t.Errorf("method should be nil: %#v", method)
	}
}

func TestClassGetClassMethod(t *testing.T) {
	className := "ClassWithClassMethod"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)

	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), "v@:")

	if method := Class_getClassMethod(class, sel); method == nil {
		t.Error("method should not be nil")
	}
}

func TestClassGetEmptyClassMethod(t *testing.T) {
	className := "ClassWithoutClassMethod"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	sel := Sel_registerName(methodeName)

	if method := Class_getClassMethod(class, sel); method != nil {
		t.Errorf("method should be nil: %#v", method)
	}
}

func TestClassCopyMethodList(t *testing.T) {
	className := "ClassWithMethods"
	class := Objc_allocateClassPair(nil, className, 0)

	closureA := ffi.Closure(func(id Id, sel Sel) {})
	closureB := ffi.Closure(func(id Id, sel Sel, n int) {})
	selA := Sel_registerName("MethodA")
	selB := Sel_registerName("MethodB")

	Class_addMethod(class, selA, Imp(unsafe.Pointer(closureA.Pointer())), "v@:")
	Class_addMethod(class, selB, Imp(unsafe.Pointer(closureB.Pointer())), "v@:i")

	methods := Class_copyMethodList(class)

	if l := len(methods); l != 2 {
		t.Errorf("methods len should be 2: %d", l)
	}
}

func TestClassCopyEmptyMethodList(t *testing.T) {
	className := "ClassWithoutMethods"
	class := Objc_allocateClassPair(nil, className, 0)

	methods := Class_copyMethodList(class)

	if l := len(methods); l != 0 {
		t.Errorf("methods len should be 0: %d", l)
	}
}

func TestClassReplaceMethod(t *testing.T) {
	className := "ClassWithMethodToBeReplaced"
	class := Objc_allocateClassPair(nil, className, 0)

	closureA := ffi.Closure(func(id Id, sel Sel) {})
	closureB := ffi.Closure(func(id Id, sel Sel) {})
	impA := Imp(unsafe.Pointer(closureA.Pointer()))
	impB := Imp(unsafe.Pointer(closureB.Pointer()))
	sel := Sel_registerName("MethodToBeReplaced")

	Class_addMethod(class, sel, impA, "v@:")

	if retImp := Class_replaceMethod(class, sel, impB, "v@:"); retImp != impA {
		t.Errorf("retImp is not the old implementation: %p != %p", retImp, impA)
	}
}

func TestClassReplaceNonexistentMethod(t *testing.T) {
	className := "ClassWithoutMethod"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)
	imp := Imp(unsafe.Pointer(closure.Pointer()))

	if retImp := Class_replaceMethod(class, sel, imp, "v@:"); retImp != nil {
		t.Errorf("retImp should be nil: %#v", retImp)
	}
}

func TestClassGetMethodImplementation(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethod"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodeName := "MethodA:"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)
	imp := Imp(unsafe.Pointer(closure.Pointer()))
	Class_addMethod(class, sel, imp, "v@:")

	if retImp := Class_getMethodImplementation(class, sel); retImp != imp {
		t.Errorf("retImp should be %p: %p", imp, retImp)
	}
}

func TestClassGetMethodImplementationStret(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethod"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodeName := "MethodA:"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)
	imp := Imp(unsafe.Pointer(closure.Pointer()))
	Class_addMethod(class, sel, imp, "v@:")

	if retImp := Class_getMethodImplementation_stret(class, sel); retImp != imp {
		t.Errorf("retImp should be %p: %p", imp, retImp)
	}
}

func TestClassRespondsToSelector(t *testing.T) {
	className := "ClassRespondingSel"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA:"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodeName)
	imp := Imp(unsafe.Pointer(closure.Pointer()))
	Class_addMethod(class, sel, imp, "v@:")

	if !Class_respondsToSelector(class, sel) {
		t.Errorf("class should responds to selector %s", methodeName)
	}
}

func TestClassNotRespondsToSelector(t *testing.T) {
	className := "ClassRespondingSel"
	class := Objc_allocateClassPair(nil, className, 0)

	methodeName := "MethodA:"
	sel := Sel_registerName(methodeName)

	if Class_respondsToSelector(class, sel) {
		t.Errorf("class should not responds to selector %s", methodeName)
	}
}

func TestClassAddProtocol(t *testing.T) {
	className := "ClassWithProtocol"
	class := Objc_allocateClassPair(nil, className, 0)
	nsObjectProto := Objc_getProtocol("NSObject")

	if !Class_addProtocol(class, nsObjectProto) {
		t.Errorf("failed to add NSObject protocol to %s", className)
	}
}

func TestClassAddAddedProtocol(t *testing.T) {
	className := "ClassWithProtocol"
	class := Objc_allocateClassPair(nil, className, 0)
	nsObjectProto := Objc_getProtocol("NSObject")
	Class_addProtocol(class, nsObjectProto)

	if Class_addProtocol(class, nsObjectProto) {
		t.Errorf("add NSObjectProtocol to %s should have failed", className)
	}
}

func TestClassAddProperty(t *testing.T) {
	className := "ClassWithProperty"
	class := Objc_allocateClassPair(nil, className, 0)

	propertyName := "A"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	if !Class_addProperty(class, propertyName, attributes) {
		t.Errorf("failed to add property %s to class %s", propertyName, className)
	}
}

func TestClassAddAddedProperty(t *testing.T) {
	className := "ClassWithProperty"
	class := Objc_allocateClassPair(nil, className, 0)

	propertyName := "A"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Class_addProperty(class, propertyName, attributes)

	if Class_addProperty(class, propertyName, attributes) {
		t.Errorf("add property %s to class %s should have failed", propertyName, className)
	}
}

func TestClassReplaceProperty(t *testing.T) {
	className := "ClassWithProperty"
	class := Objc_allocateClassPair(nil, className, 0)

	propertyName := "A"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Class_addProperty(class, propertyName, attributes)

	attributes = []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "i"},
		PropertyAttribute{Name: "V"},
	}

	Class_replaceProperty(class, propertyName, attributes)

	property := Class_getProperty(class, propertyName)
	expectedAttributeString := "Ti,V"

	if attributesString := Property_getAttributes(property); attributesString != expectedAttributeString {
		t.Errorf("attributes string should be %s: %s", expectedAttributeString, attributesString)
	}
}

func TestClassConformToProtocol(t *testing.T) {
	className := "ClassWithProtocol"
	class := Objc_allocateClassPair(nil, className, 0)
	nsObjectProto := Objc_getProtocol("NSObject")

	Class_addProtocol(class, nsObjectProto)

	if !Class_conformsToProtocol(class, nsObjectProto) {
		t.Error("class should conform to NSObject protocol")
	}
}

func TestClassNotConformToProtocol(t *testing.T) {
	className := "ClassWithoutProtocol"
	class := Objc_allocateClassPair(nil, className, 0)
	nsObjectProto := Objc_getProtocol("NSObject")

	if Class_conformsToProtocol(class, nsObjectProto) {
		t.Error("class should not conform to NSObject protocol")
	}
}

func TestClassCopyProtocolList(t *testing.T) {
	className := "ClassWithProtocol"
	class := Objc_allocateClassPair(nil, className, 0)
	nsObjectProto := Objc_getProtocol("NSObject")

	Class_addProtocol(class, nsObjectProto)

	protocols := Class_copyProtocolList(class)

	if l := len(protocols); l != 1 {
		t.Errorf("protocols len should be 1: %d", l)
	}
}

func TestClassCopyEmptyProtocolList(t *testing.T) {
	className := "ClassWithoutProtocol"
	class := Objc_allocateClassPair(nil, className, 0)

	protocols := Class_copyProtocolList(class)

	if l := len(protocols); l != 0 {
		t.Errorf("protocols len should be 0: %d", l)
	}
}

func TestClassGetVersion(t *testing.T) {
	className := "Class"
	class := Objc_allocateClassPair(nil, className, 0)

	if version := Class_getVersion(class); version != 0 {
		t.Errorf("class version should be 0: %d", version)
	}
}

func TestClassSetVersion(t *testing.T) {
	className := "Class"
	class := Objc_allocateClassPair(nil, className, 0)

	Class_setVersion(class, 42)

	if version := Class_getVersion(class); version != 42 {
		t.Errorf("class version should be 42: %d", version)
	}
}

func TestClassCreateInstance(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	class := Objc_allocateClassPair(nsObject, "ClassToBeInstanciated", 0)
	Objc_registerClassPair(class)

	if instance := Class_createInstance(class, 0); instance == nil {
		t.Error("instance should not be nil")
	}
}

func TestClassGetImageName(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	libName := "/usr/lib/libobjc.A.dylib"

	if image := Class_getImageName(nsObject); image != libName {
		t.Errorf("image should be %s: %s", libName, image)
	}
}
