package objc

import (
	"testing"
	"unsafe"
)

func TestCallocFree(t *testing.T) {
	var ptr unsafe.Pointer

	if ptr = calloc(2, 4); ptr == nil {
		t.Error("ptr should not be nil")
	}

	free(ptr)
}
