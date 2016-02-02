package objc

import (
	"testing"
	"unsafe"

	"github.com/achille-roussel/go-ffi"
)

func TestMethodGetName(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	sel := Sel_registerName("methodForSelector:")
	method := Class_getInstanceMethod(nsObject, sel)

	if methodName := Method_getName(method); methodName != sel {
		t.Errorf("methodName should be %#v: %#v", sel, methodName)
	}
}

func TestMethodGetImplementation(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	sel := Sel_registerName("methodForSelector:")
	method := Class_getInstanceMethod(nsObject, sel)

	if imp := Method_getImplementation(method); imp == nil {
		t.Error("methodName should not be nil")
	}
}

func TestMethodGetTypeEncoding(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethodForTypeTest"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodName)
	methodTypes := "v@:"
	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), methodTypes)

	method := Class_getInstanceMethod(class, sel)

	if types := Method_getTypeEncoding(method); types != methodTypes {
		t.Errorf("types should be %s: %s", methodTypes, types)
	}
}

func TestMethodCopyReturnType(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethodForRetTest"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodName)
	methodTypes := "v@:"
	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), methodTypes)

	method := Class_getInstanceMethod(class, sel)
	expectedReturnType := "v"

	if returnType := Method_copyReturnType(method); returnType != expectedReturnType {
		t.Errorf("returnType should be %s: %s", expectedReturnType, returnType)
	}
}

func TestMethodCopyArgumentType(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethodForArgTest"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodName)
	methodTypes := "v@:"
	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), methodTypes)

	method := Class_getInstanceMethod(class, sel)
	expectedArgType := "@"

	if argType := Method_copyArgumentType(method, 0); argType != expectedArgType {
		t.Errorf("argTypes should be %s: %s", expectedArgType, argType)
	}
}

func TestMethodGetNumberOfArguments(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethodForArgCountTest"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodName)
	methodTypes := "v@:"
	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), methodTypes)

	method := Class_getInstanceMethod(class, sel)
	expectedArgCount := uint(2)

	if count := Method_getNumberOfArguments(method); count != expectedArgCount {
		t.Errorf("argTypes should be %d: %d", expectedArgCount, count)
	}
}

func TestMethodGetDescription(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassWithMethodForDescriptionTest"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodName := "MethodA"
	closure := ffi.Closure(func(id Id, sel Sel) {})
	sel := Sel_registerName(methodName)
	methodTypes := "v@:"
	Class_addMethod(class, sel, Imp(unsafe.Pointer(closure.Pointer())), methodTypes)

	method := Class_getInstanceMethod(class, sel)
	description := Method_getDescription(method)

	if description.Name != sel {
		t.Errorf("description name should be %#v: %#v", sel, description.Name)
	}

	if description.Types != methodTypes {
		t.Errorf("description types should be %s: %s", methodTypes, description.Types)
	}
}

func TestMethodSetImplementation(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassForExchangeMethodImp"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodName := "MethodA"
	closureA := ffi.Closure(func(id Id, sel Sel) {})
	impA := Imp(unsafe.Pointer(closureA.Pointer()))
	closureB := ffi.Closure(func(id Id, sel Sel) {})
	impB := Imp(unsafe.Pointer(closureB.Pointer()))
	sel := Sel_registerName(methodName)
	methodTypes := "v@:"
	Class_addMethod(class, sel, impA, methodTypes)

	method := Class_getInstanceMethod(class, sel)

	if oldImp := Method_setImplementation(method, impB); oldImp != impA {
		t.Errorf("old imp should be %p: %p", impA, oldImp)
	}
}

func TestMethodExchangeImplementations(t *testing.T) {
	nsObject := Objc_getClass("NSObject")
	className := "ClassForExchangeMethodImp"
	class := Objc_allocateClassPair(nsObject, className, 0)

	methodAName := "MethodA"
	closureA := ffi.Closure(func(id Id, sel Sel) {})
	impA := Imp(unsafe.Pointer(closureA.Pointer()))
	selA := Sel_registerName(methodAName)
	methodTypes := "v@:"
	Class_addMethod(class, selA, impA, methodTypes)
	methodA := Class_getInstanceMethod(class, selA)

	methodBName := "MethodB"
	closureB := ffi.Closure(func(id Id, sel Sel) {})
	impB := Imp(unsafe.Pointer(closureB.Pointer()))
	selB := Sel_registerName(methodBName)
	Class_addMethod(class, selB, impB, methodTypes)
	methodB := Class_getInstanceMethod(class, selB)

	Method_exchangeImplementations(methodA, methodB)

	if newImpA := Method_getImplementation(methodA); newImpA != impB {
		t.Errorf("newImpA should be %p: %p", impB, newImpA)
	}

	if newImpB := Method_getImplementation(methodB); newImpB != impA {
		t.Errorf("newImpB should be %p: %p", impA, newImpB)
	}
}
