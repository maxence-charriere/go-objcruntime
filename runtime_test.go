package objc

import "testing"

func TestSelEquality(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selABis := Sel_registerName("SelectorA")
	selB := Sel_registerName("SelectorB")

	if !Sel_isEqual(selA, selABis) {
		t.Errorf("selA: %p and selABis: %p should be equal", selA, selABis)
	}

	if Sel_isEqual(selA, selB) {
		t.Errorf("selA: %p and selB: %p should not be equal", selA, selB)
	}
}

func TestSelName(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selAName := Sel_getName(selA)
	selABis := Sel_registerName("SelectorA")
	selABisName := Sel_getName(selABis)
	selB := Sel_registerName("SelectorB")
	selBName := Sel_getName(selB)

	if selAName != selABisName {
		t.Errorf("selAName and selBName should be equal: %s != %s\n", selAName, selABisName)
	}

	if selAName == selBName {
		t.Errorf("selAName and selBName should not be equal: %s == %s\n", selAName, selBName)
	}

}
