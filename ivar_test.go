package objc

import "testing"

func TestIvarGetName(t *testing.T) {
	className := "ClassWithIvar"
	class := Objc_allocateClassPair(nil, className, 0)
	ivarName := "ivar"

	Class_addIvar(class, ivarName, 4, 0, "i")
	ivar := Class_getInstanceVariable(class, ivarName)

	if name := Ivar_getName(ivar); name != ivarName {
		t.Errorf("name should be %s: %s", ivarName, ivar)
	}
}

func TestIvarGetTypeEncoding(t *testing.T) {
	className := "ClassWithIvarForEncodingTest"
	ivarName := "ivar"
	typeEncoding := "i"
	class := Objc_allocateClassPair(nil, className, 0)

	Class_addIvar(class, ivarName, 4, 0, typeEncoding)
	ivar := Class_getInstanceVariable(class, ivarName)

	if encoding := Ivar_getTypeEncoding(ivar); encoding != typeEncoding {
		t.Errorf("encoding should be %s: %s", typeEncoding, encoding)
	}
}
