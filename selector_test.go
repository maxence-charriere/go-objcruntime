package objc

import "testing"

func TestSelNameEquality(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selAName := Sel_getName(selA)
	selB := Sel_registerName("SelectorA")
	selBName := Sel_getName(selB)

	if selAName != selBName {
		t.Errorf("selAName and selBName should be equal: %s != %s\n", selAName, selBName)
	}
}

func TestSelNameDifference(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selAName := Sel_getName(selA)
	selB := Sel_registerName("SelectorB")
	selBName := Sel_getName(selB)

	if selAName == selBName {
		t.Errorf("selAName and selBName should not be equal: %s == %s\n", selAName, selBName)
	}
}

func TestSelEmpty(t *testing.T) {
	sel := Sel_registerName("")

	if name := Sel_getName(sel); name != "" {
		t.Errorf("sel name should be a empty string: %s\n", name)
	}
}

func TestSelEquality(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selB := Sel_registerName("SelectorA")

	if !Sel_isEqual(selA, selB) {
		t.Errorf("selA: %p and selB: %p should be equal", selA, selB)
	}
}

func TestSelDifference(t *testing.T) {
	selA := Sel_registerName("SelectorA")
	selB := Sel_registerName("SelectorB")

	if Sel_isEqual(selA, selB) {
		t.Errorf("selA: %p and selB: %p should not be equal", selA, selB)
	}
}

func TestSelGetUid(t *testing.T) {
	sel := Sel_getUid("Selector")
	selBis := Sel_registerName("Selector")

	if !Sel_isEqual(sel, selBis) {
		t.Errorf("sel: %p and selBis: %p should be equal", sel, selBis)
	}
}
