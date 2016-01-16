package objc

import "testing"

func TestCalloc(t *testing.T) {
	if ptr := calloc(2, 4); ptr == nil {
		t.Error("ptr should not be nil")
	}
}
