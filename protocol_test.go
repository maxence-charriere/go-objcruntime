package objc

import "testing"

func TestGetProtocol(t *testing.T) {
	if p := Objc_getProtocol("NSObject"); p == nil {
		t.Error("protocol should not be nil")
	}
}

func TestGetNonExistentProtocol(t *testing.T) {
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

func TestProtocolAddMethodDescription(t *testing.T) {
	proto := Objc_allocateProtocol("ProtoWithMethod")
	sel := Sel_registerName("TestA")
	types := "i^c"

	Protocol_addMethodDescription(proto, sel, types, false, false)
}

func TestProtocolAddMethodDescriptionWithoutType(t *testing.T) {
	proto := Objc_allocateProtocol("ProtoWithMethodWithoutType")
	sel := Sel_registerName("TestA")

	Protocol_addMethodDescription(proto, sel, "", false, false)
}

func TestProtocolAddProtocol(t *testing.T) {
	proto := Objc_allocateProtocol("CompositionProto")
	nsObjectProto := Objc_getProtocol("NSObject")

	Protocol_addProtocol(proto, nsObjectProto)
	Objc_registerProtocol(proto)
}

func TestProtocolAddProperty(t *testing.T) {
	var attributes []PropertyAttribute

	proto := Objc_allocateProtocol("ProtoWithProperty")
	name := "Property"

	Protocol_addProperty(proto, name, attributes, true, true)
	Objc_registerProtocol(proto)
}

func TestProtocolAddPropertyWithAttributes(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyWithAttributesProto")
	name := "CharDefault"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Protocol_addProperty(proto, name, attributes, true, true)
	Objc_registerProtocol(proto)
}

func TestProtocolGetName(t *testing.T) {
	name := "NSObject"
	proto := Objc_getProtocol(name)

	if protoName := Protocol_getName(proto); protoName != name {
		t.Errorf("protocol name should be %s: %s", name, protoName)
	}
}

func TestProtocolEquality(t *testing.T) {
	protoA := Objc_getProtocol("NSObject")
	protoB := Objc_getProtocol("NSObject")

	if !Protocol_isEqual(protoA, protoB) {
		t.Errorf("protoA and protoB should be equal: %p != %p", protoA, protoB)
	}
}

func TestProtocolDifference(t *testing.T) {
	protoA := Objc_allocateProtocol("ProtocolDiffA")
	protoB := Objc_allocateProtocol("ProtocolDiffB")

	Objc_registerProtocol(protoA)
	Objc_registerProtocol(protoB)

	if Protocol_isEqual(protoA, protoB) {
		t.Errorf("protoA and protoB should not be equal: %p != %p", protoA, protoB)
	}
}

func TestProtocolCopyMethodDescriptionList(t *testing.T) {
	proto := Objc_allocateProtocol("CopyMethodDescriptionListProto")
	selName := "TestA"
	sel := Sel_registerName(selName)
	types := "i^c"
	Protocol_addMethodDescription(proto, sel, types, false, false)
	Objc_registerProtocol(proto)

	descriptions := Protocol_copyMethodDescriptionList(proto, false, false)

	if l := len(descriptions); l != 1 {
		t.Fatalf("desciptions should have 1 element: %d", l)
	}

	description := descriptions[0]

	if descriptionName := Sel_getName(description.Name); descriptionName != selName {
		t.Errorf("description name should be named %s: %s", selName, descriptionName)
	}

	if description.Types != types {
		t.Errorf("description types should be %s: %s", types, description.Types)
	}
}

func TestProtocolCopyNonConformMethodDescriptionList(t *testing.T) {
	proto := Objc_allocateProtocol("CopyNonConformMethodDescriptionListProto")
	selName := "TestA"
	sel := Sel_registerName(selName)
	types := "i^c"
	Protocol_addMethodDescription(proto, sel, types, false, false)
	Objc_registerProtocol(proto)

	descriptions := Protocol_copyMethodDescriptionList(proto, false, true)

	if l := len(descriptions); l != 0 {
		t.Errorf("desciptions should be empty: %d", l)
	}
}

func TestProtocolGetMethodDescription(t *testing.T) {
	proto := Objc_allocateProtocol("GetMethodDescriptionProto")
	selName := "TestA"
	sel := Sel_registerName(selName)
	types := "i^c"
	Protocol_addMethodDescription(proto, sel, types, true, true)
	Objc_registerProtocol(proto)

	description := Protocol_getMethodDescription(proto, sel, true, true)

	if descriptionName := Sel_getName(description.Name); descriptionName != selName {
		t.Errorf("description Sel name should be named %s: %s", selName, descriptionName)
	}

	if description.Types != types {
		t.Errorf("description types should be %s: %s", types, description.Types)
	}
}

func TestProtocolGetNonConformMethodDescription(t *testing.T) {
	proto := Objc_allocateProtocol("GetNonConformMethodDescriptionProto")
	selName := "TestA"
	sel := Sel_registerName(selName)
	types := "i^c"
	Protocol_addMethodDescription(proto, sel, types, true, true)
	Objc_registerProtocol(proto)

	description := Protocol_getMethodDescription(proto, sel, false, true)

	if description.Name != nil {
		t.Errorf("description name should be nil: %#v", description.Name)
	}

	if description.Types != "" {
		t.Errorf("description types should be an empty string: %s", description.Types)
	}
}

func TestProtocolCopyMethodDescriptionListNonConform(t *testing.T) {
	proto := Objc_allocateProtocol("CopyMethodNonConformProto")
	selName := "TestA"
	sel := Sel_registerName(selName)
	types := "i^c"
	Protocol_addMethodDescription(proto, sel, types, false, false)
	Objc_registerProtocol(proto)

	descriptions := Protocol_copyMethodDescriptionList(proto, false, true)

	if l := len(descriptions); l != 0 {
		t.Fatalf("desciptions should have 0 elements: %d", l)
	}
}

func TestProtocolCopyPropertyList(t *testing.T) {
	proto := Objc_allocateProtocol("CopyPropertyListProto")
	propertyName := "Property"

	Protocol_addProperty(proto, propertyName, nil, true, true)
	Objc_registerProtocol(proto)

	properties := Protocol_copyPropertyList(proto)

	if l := len(properties); l != 1 {
		t.Fatalf("properties should have 1 element: %d", l)
	}

	if name := Propety_getName(properties[0]); name != propertyName {
		t.Errorf("property name should be %s: %s", propertyName, name)
	}
}

func TestProtocolCopyPropertyListEmpty(t *testing.T) {
	proto := Objc_allocateProtocol("CopyPropertyListEmptyProto")
	Objc_registerProtocol(proto)

	properties := Protocol_copyPropertyList(proto)

	if l := len(properties); l != 0 {
		t.Fatalf("properties len should be empty: %d", l)
	}
}

func TestProtocolGetProperty(t *testing.T) {
	proto := Objc_allocateProtocol("GetPropertyProto")
	propertyName := "CharDefault"

	Protocol_addProperty(proto, propertyName, nil, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)

	if property == nil {
		t.Fatal("property should not be nil")
	}

	if name := Propety_getName(property); name != propertyName {
		t.Errorf("property name should be %s: %s", propertyName, name)
	}
}

func TestProtocolGetNonExistentProperty(t *testing.T) {
	proto := Objc_allocateProtocol("GetPropertyEmptyProto")
	Objc_registerProtocol(proto)

	if property := Protocol_getProperty(proto, "Property", true, true); property != nil {
		t.Errorf("property should be nil: %#v", property)
	}
}

func TestProtocolGetNonConformProperty(t *testing.T) {
	proto := Objc_allocateProtocol("GetNonConformPropertyProto")
	propertyName := "Property"

	Protocol_addProperty(proto, propertyName, nil, true, true)
	Objc_registerProtocol(proto)

	if property := Protocol_getProperty(proto, propertyName, false, true); property != nil {
		t.Errorf("property should be nil: %#v", property)
	}
}

func TestProtocolCopyProtocolList(t *testing.T) {
	proto := Objc_allocateProtocol("CopyProtocolListProto")
	nsObjectProto := Objc_getProtocol("NSObject")
	Protocol_addProtocol(proto, nsObjectProto)
	Objc_registerProtocol(proto)

	protocols := Protocol_copyProtocolList(proto)

	if l := len(protocols); l != 1 {
		t.Fatalf("protocols should have 1 element: %d", l)
	}

	if !Protocol_isEqual(protocols[0], nsObjectProto) {
		t.Errorf("Protocol is not an NSObject: %#v != %#v", protocols[0], nsObjectProto)
	}
}

func TestProtocolCopyProtocolListEmpty(t *testing.T) {
	proto := Objc_allocateProtocol("CopyProtocolListEmptyProto")
	Objc_registerProtocol(proto)

	protocols := Protocol_copyProtocolList(proto)

	if l := len(protocols); l != 0 {
		t.Fatalf("protocols should be empty: %d", l)
	}
}

func TestProtocolConformToProtocol(t *testing.T) {
	proto := Objc_allocateProtocol("ProtoConformToProto")
	nsObjectProto := Objc_getProtocol("NSObject")
	Protocol_addProtocol(proto, nsObjectProto)
	Objc_registerProtocol(proto)

	if !Protocol_conformsToProtocol(proto, nsObjectProto) {
		t.Error("protocol should conform to NSOject")
	}
}

func TestProtocolNotConformToProtocol(t *testing.T) {
	proto := Objc_allocateProtocol("ProtoNotConformToProto")
	nsObjectProto := Objc_getProtocol("NSObject")
	Objc_registerProtocol(proto)

	if Protocol_conformsToProtocol(proto, nsObjectProto) {
		t.Error("protocol should not conform to NSOject")
	}
}
