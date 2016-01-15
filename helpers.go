package objc

// #include <objc/runtime.h>
import "C"

func cBool(value bool) C.BOOL {
	if value {
		return 1
	}

	return 0
}
