package objcruntime

// #include <objc/runtime.h>
import "C"

// Ivar define an opaque type that represents an instance variable.
type Ivar C.Ivar
