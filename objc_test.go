package objc

import "testing"

func TestAllocateClassPair(t *testing.T) {
	if class := Objc_allocateClassPair(nil, "ClassToAllocate", 0); class == nil {
		t.Error("class should not be nil")
	}
}

func TestAllocateClassPairWithSuperClass(t *testing.T) {
	superclass := Objc_getClass("NSObject")

	if class := Objc_allocateClassPair(superclass, "ClassToAllocateWithSuper", 0); class == nil {
		t.Error("class should not be nil")
	}
}

func TestAllocateExistingClass(t *testing.T) {
	className := "NSObject"

	if class := Objc_allocateClassPair(nil, className, 0); class != nil {
		t.Errorf("class should be nil: %#v", class)
	}
}

func TestDisposeClassPair(t *testing.T) {
	superclass := Objc_getClass("NSObject")
	class := Objc_allocateClassPair(superclass, "ClassToDispose", 0)
	Objc_disposeClassPair(class)
}

func TestRegisterClassPair(t *testing.T) {
	class := Objc_allocateClassPair(nil, "ClassToRegister", 0)
	Objc_registerClassPair(class)
}

func TestConstructInstance(t *testing.T) {
	superclass := Objc_getClass("NSObject")
	class := Objc_allocateClassPair(superclass, "ClassToConstruct", 0)
	Objc_registerClassPair(class)

	bytes := calloc(1, uintptr(Class_getInstanceSize(class)))

	if instance := Objc_constructInstance(class, bytes); instance == nil {
		t.Error("instance should not be nil")
	}
}

func TestConstructNilInstance(t *testing.T) {
	superclass := Objc_getClass("NSObject")
	class := Objc_allocateClassPair(superclass, "ClassToConstructNil", 0)
	Objc_registerClassPair(class)

	if instance := Objc_constructInstance(class, nil); instance != nil {
		t.Errorf("instance should be nil: %#v", instance)
	}
}

func TestDestructInstance(t *testing.T) {
	superclass := Objc_getClass("NSObject")
	class := Objc_allocateClassPair(superclass, "ClassToDestruct", 0)
	Objc_registerClassPair(class)

	instance := Class_createInstance(class, 0)
	Objc_destructInstance(instance)

}

func TestDestructNil(t *testing.T) {
	Objc_destructInstance(nil)
}

func TestCopyClassList(t *testing.T) {
	classes := Objc_copyClassList()

	if len(classes) == 0 {
		t.Error("Classes len should not be 0")
	}
}

func TestGetClass(t *testing.T) {
	if class := Objc_getClass("NSObject"); class == nil {
		t.Error("class should not be nil")
	}
}

func TestGetNonexistentClassClass(t *testing.T) {
	if class := Objc_getClass("NonexistentNSObject"); class != nil {
		t.Errorf("class should be nil: %#v", class)
	}
}

func TestGetMetaClass(t *testing.T) {
	if metaclass := Objc_getMetaClass("NSObject"); metaclass == nil {
		t.Error("metaclass should not be nil")
	}
}

func TesGettNonexistentMetaClass(t *testing.T) {
	if metaclass := Objc_getMetaClass("NonexistentNSObject"); metaclass != nil {
		t.Error("metaclass should be nil: %#v", metaclass)
	}
}

func TestGetProtocol(t *testing.T) {
	if p := Objc_getProtocol("NSObject"); p == nil {
		t.Error("protocol should not be nil")
	}
}

func TestGetNonexistentProtocol(t *testing.T) {
	if p := Objc_getProtocol("MyCustomProtocol"); p != nil {
		t.Errorf("protocol should be nil:%#v", p)
	}
}

func TestCopyProtocolList(t *testing.T) {
	if protocols := Objc_copyProtocolList(); len(protocols) == 0 {
		t.Error("protocol list should not be empty")
	}
}

func TestAllocateProtocol(t *testing.T) {
	if p := Objc_allocateProtocol("AllocatedProtocol"); p == nil {
		t.Error("allocated protocol should not be nil")
	}
}

func TestAllocateExistentProtocol(t *testing.T) {
	if p := Objc_allocateProtocol("NSObject"); p != nil {
		t.Errorf("allocated protocol should be nil: %#v", p)
	}
}

func TestRegisterProtocol(t *testing.T) {
	proto := Objc_allocateProtocol("RegisteredProto")
	Objc_registerProtocol(proto)
	protocols := Objc_copyProtocolList()

	for _, p := range protocols {
		if p == proto {
			return
		}
	}

	t.Errorf("Registered protocol not found: %#v", protocols)
}

func TestFindNotRegisteredProtocol(t *testing.T) {
	proto := Objc_allocateProtocol("NotRegisteredProto")
	protocols := Objc_copyProtocolList()

	for _, p := range protocols {
		if p == proto {
			t.Fatalf("GoTestable protocol should not be registered: %#v", p)
		}
	}
}
