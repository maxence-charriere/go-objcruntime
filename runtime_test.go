package objc

import "testing"

func TestSel(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selAName := Sel_getName(selA)
	selB := Sel_getUid("SelectorA")

	if !Sel_isEqual(selA, selB) {
		t.Errorf("selA: %p and selB: %p should be equal", selA, selB)
	}

	if selBName := Sel_getName(selB); selAName != selBName {
		t.Errorf("selAName and selBName are not equal: %s != %s\n", selAName, selBName)
	}
}
