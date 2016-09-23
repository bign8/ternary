// This file has been generated with: go run gen.go; DO NOT EDIT!

package ternary

import "unsafe"

// Bool provides the ternary operator for bool types
func Bool(ok bool, truthy, falsey bool) bool {
	if ok {
		return truthy
	}
	return falsey
}

// Complex128 provides the ternary operator for complex128 types
func Complex128(ok bool, truthy, falsey complex128) complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// Complex64 provides the ternary operator for complex64 types
func Complex64(ok bool, truthy, falsey complex64) complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// Float32 provides the ternary operator for float32 types
func Float32(ok bool, truthy, falsey float32) float32 {
	if ok {
		return truthy
	}
	return falsey
}

// Float64 provides the ternary operator for float64 types
func Float64(ok bool, truthy, falsey float64) float64 {
	if ok {
		return truthy
	}
	return falsey
}

// Int provides the ternary operator for int types
func Int(ok bool, truthy, falsey int) int {
	if ok {
		return truthy
	}
	return falsey
}

// Int16 provides the ternary operator for int16 types
func Int16(ok bool, truthy, falsey int16) int16 {
	if ok {
		return truthy
	}
	return falsey
}

// Int32 provides the ternary operator for int32 types
func Int32(ok bool, truthy, falsey int32) int32 {
	if ok {
		return truthy
	}
	return falsey
}

// Int64 provides the ternary operator for int64 types
func Int64(ok bool, truthy, falsey int64) int64 {
	if ok {
		return truthy
	}
	return falsey
}

// Int8 provides the ternary operator for int8 types
func Int8(ok bool, truthy, falsey int8) int8 {
	if ok {
		return truthy
	}
	return falsey
}

// Interface provides the ternary operator for interface{} types
func Interface(ok bool, truthy, falsey interface{}) interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// String provides the ternary operator for string types
func String(ok bool, truthy, falsey string) string {
	if ok {
		return truthy
	}
	return falsey
}

// Uint provides the ternary operator for uint types
func Uint(ok bool, truthy, falsey uint) uint {
	if ok {
		return truthy
	}
	return falsey
}

// Uint16 provides the ternary operator for uint16 types
func Uint16(ok bool, truthy, falsey uint16) uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// Uint32 provides the ternary operator for uint32 types
func Uint32(ok bool, truthy, falsey uint32) uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// Uint64 provides the ternary operator for uint64 types
func Uint64(ok bool, truthy, falsey uint64) uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// Uint8 provides the ternary operator for uint8 types
func Uint8(ok bool, truthy, falsey uint8) uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// Uintptr provides the ternary operator for uintptr types
func Uintptr(ok bool, truthy, falsey uintptr) uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// UnsafePointer provides the ternary operator for unsafe.Pointer types
func UnsafePointer(ok bool, truthy, falsey unsafe.Pointer) unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}
