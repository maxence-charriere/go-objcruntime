package objc

import "testing"

func TestMakePropertyAttribute(t *testing.T) {
	attr := PropertyAttribute{
		Name:  "T",
		Value: "c",
	}

	cattr := attr.ctype()
	newAttr := makePropertyAttribute(cattr)

	if attr.Name != newAttr.Name {
		t.Errorf("attribute name should be %s: %s", attr.Name, newAttr.Name)
	}

	if attr.Value != newAttr.Value {
		t.Errorf("attribute value should be %s: %s", attr.Value, newAttr.Value)
	}
}

func TestPropertyGetName(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyNameProto")
	propertyName := "CharDefault"

	Protocol_addProperty(proto, propertyName, nil, true, true)
	Objc_registerProtocol(proto)
	property := Protocol_getProperty(proto, propertyName, true, true)

	if name := Propety_getName(property); name != propertyName {
		t.Errorf("property name should be %s: %s", propertyName, name)
	}
}

func TestPropertyGetAttributes(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyAttributesProto")
	propertyName := "CharDefault"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Protocol_addProperty(proto, propertyName, attributes, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)
	expectedAttributes := "Tc,VcharDefault"

	if attibutesString := Property_getAttributes(property); attibutesString != expectedAttributes {
		t.Errorf("attributes should be %s:%s", expectedAttributes, attibutesString)
	}
}

func TestPropertyGetAttributesEmpty(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyAttributesEmptyProto")
	propertyName := "CharDefault"

	Protocol_addProperty(proto, propertyName, nil, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)

	if attibutesString := Property_getAttributes(property); len(attibutesString) != 0 {
		t.Errorf("attributes should be empty: %s", attibutesString)
	}
}

func TestPropertyCopyAttributeValue(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyCopyAttrValProto")
	propertyName := "CharDefault"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Protocol_addProperty(proto, propertyName, attributes, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)

	if value := Property_copyAttributeValue(property, "T"); value != "c" {
		t.Errorf("value should be c: %s", value)
	}
}

func TestPropertyCopyAttributeEmptyValue(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyCopyAttrValEmptyProto")
	propertyName := "CharDefault"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V"},
	}

	Protocol_addProperty(proto, propertyName, attributes, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)

	if value := Property_copyAttributeValue(property, "V"); value != "" {
		t.Error("value should be an empty string: %s", value)
	}
}

func TestPropertyCopyNonExistentAttributeValue(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyCopyNonExistAttrValProto")
	propertyName := "CharDefault"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
		PropertyAttribute{Name: "V", Value: "charDefault"},
	}

	Protocol_addProperty(proto, propertyName, attributes, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)

	if value := Property_copyAttributeValue(property, "42"); value != "" {
		t.Error("value should be an empty string: %s", value)
	}
}

func TestPropertyCopyAttributeList(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyCopyAttributesProto")
	propertyName := "CharDefault"

	attributes := []PropertyAttribute{
		PropertyAttribute{Name: "T", Value: "c"},
	}

	Protocol_addProperty(proto, propertyName, attributes, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)
	attributes = Property_copyAttributeList(property)

	if l := len(attributes); l != 1 {
		t.Fatal("attibutes len should be 1: %d", l)
	}

	if attrName := attributes[0].Name; attrName != "T" {
		t.Error("attribute name should be T:", attrName)
	}

	if attrVal := attributes[0].Value; attrVal != "c" {
		t.Error("attribute value should be c:", attrVal)
	}
}

func TestPropertyCopyAttributeListEmpty(t *testing.T) {
	proto := Objc_allocateProtocol("PropertyCopyAttributesEmptyProto")
	propertyName := "CharDefault"

	Protocol_addProperty(proto, propertyName, nil, true, true)
	Objc_registerProtocol(proto)

	property := Protocol_getProperty(proto, propertyName, true, true)
	attributes := Property_copyAttributeList(property)

	if l := len(attributes); l != 0 {
		t.Fatalf("attibutes len should be 0:%d", l)
	}
}
