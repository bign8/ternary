// This file has been generated with: go run gen.go; DO NOT EDIT!

package ternary

import "unsafe"

// ArrayBool provides the ternary operator for []bool types
func ArrayBool(ok bool, truthy, falsey []bool) []bool {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayComplex128 provides the ternary operator for []complex128 types
func ArrayComplex128(ok bool, truthy, falsey []complex128) []complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayComplex64 provides the ternary operator for []complex64 types
func ArrayComplex64(ok bool, truthy, falsey []complex64) []complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayFloat32 provides the ternary operator for []float32 types
func ArrayFloat32(ok bool, truthy, falsey []float32) []float32 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayFloat64 provides the ternary operator for []float64 types
func ArrayFloat64(ok bool, truthy, falsey []float64) []float64 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayInt provides the ternary operator for []int types
func ArrayInt(ok bool, truthy, falsey []int) []int {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayInt16 provides the ternary operator for []int16 types
func ArrayInt16(ok bool, truthy, falsey []int16) []int16 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayInt32 provides the ternary operator for []int32 types
func ArrayInt32(ok bool, truthy, falsey []int32) []int32 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayInt64 provides the ternary operator for []int64 types
func ArrayInt64(ok bool, truthy, falsey []int64) []int64 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayInt8 provides the ternary operator for []int8 types
func ArrayInt8(ok bool, truthy, falsey []int8) []int8 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayInterface provides the ternary operator for []interface{} types
func ArrayInterface(ok bool, truthy, falsey []interface{}) []interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayString provides the ternary operator for []string types
func ArrayString(ok bool, truthy, falsey []string) []string {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUint provides the ternary operator for []uint types
func ArrayUint(ok bool, truthy, falsey []uint) []uint {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUint16 provides the ternary operator for []uint16 types
func ArrayUint16(ok bool, truthy, falsey []uint16) []uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUint32 provides the ternary operator for []uint32 types
func ArrayUint32(ok bool, truthy, falsey []uint32) []uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUint64 provides the ternary operator for []uint64 types
func ArrayUint64(ok bool, truthy, falsey []uint64) []uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUint8 provides the ternary operator for []uint8 types
func ArrayUint8(ok bool, truthy, falsey []uint8) []uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUintptr provides the ternary operator for []uintptr types
func ArrayUintptr(ok bool, truthy, falsey []uintptr) []uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// ArrayUnsafePointer provides the ternary operator for []unsafe.Pointer types
func ArrayUnsafePointer(ok bool, truthy, falsey []unsafe.Pointer) []unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// Bool provides the ternary operator for bool types
func Bool(ok bool, truthy, falsey bool) bool {
	if ok {
		return truthy
	}
	return falsey
}

// ChanBool provides the ternary operator for chan bool types
func ChanBool(ok bool, truthy, falsey chan bool) chan bool {
	if ok {
		return truthy
	}
	return falsey
}

// ChanComplex128 provides the ternary operator for chan complex128 types
func ChanComplex128(ok bool, truthy, falsey chan complex128) chan complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanComplex64 provides the ternary operator for chan complex64 types
func ChanComplex64(ok bool, truthy, falsey chan complex64) chan complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanFloat32 provides the ternary operator for chan float32 types
func ChanFloat32(ok bool, truthy, falsey chan float32) chan float32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanFloat64 provides the ternary operator for chan float64 types
func ChanFloat64(ok bool, truthy, falsey chan float64) chan float64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanInt provides the ternary operator for chan int types
func ChanInt(ok bool, truthy, falsey chan int) chan int {
	if ok {
		return truthy
	}
	return falsey
}

// ChanInt16 provides the ternary operator for chan int16 types
func ChanInt16(ok bool, truthy, falsey chan int16) chan int16 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanInt32 provides the ternary operator for chan int32 types
func ChanInt32(ok bool, truthy, falsey chan int32) chan int32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanInt64 provides the ternary operator for chan int64 types
func ChanInt64(ok bool, truthy, falsey chan int64) chan int64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanInt8 provides the ternary operator for chan int8 types
func ChanInt8(ok bool, truthy, falsey chan int8) chan int8 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanInterface provides the ternary operator for chan interface{} types
func ChanInterface(ok bool, truthy, falsey chan interface{}) chan interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvBool provides the ternary operator for <-chan bool types
func ChanRecvBool(ok bool, truthy, falsey <-chan bool) <-chan bool {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvComplex128 provides the ternary operator for <-chan complex128 types
func ChanRecvComplex128(ok bool, truthy, falsey <-chan complex128) <-chan complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvComplex64 provides the ternary operator for <-chan complex64 types
func ChanRecvComplex64(ok bool, truthy, falsey <-chan complex64) <-chan complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvFloat32 provides the ternary operator for <-chan float32 types
func ChanRecvFloat32(ok bool, truthy, falsey <-chan float32) <-chan float32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvFloat64 provides the ternary operator for <-chan float64 types
func ChanRecvFloat64(ok bool, truthy, falsey <-chan float64) <-chan float64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvInt provides the ternary operator for <-chan int types
func ChanRecvInt(ok bool, truthy, falsey <-chan int) <-chan int {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvInt16 provides the ternary operator for <-chan int16 types
func ChanRecvInt16(ok bool, truthy, falsey <-chan int16) <-chan int16 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvInt32 provides the ternary operator for <-chan int32 types
func ChanRecvInt32(ok bool, truthy, falsey <-chan int32) <-chan int32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvInt64 provides the ternary operator for <-chan int64 types
func ChanRecvInt64(ok bool, truthy, falsey <-chan int64) <-chan int64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvInt8 provides the ternary operator for <-chan int8 types
func ChanRecvInt8(ok bool, truthy, falsey <-chan int8) <-chan int8 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvInterface provides the ternary operator for <-chan interface{} types
func ChanRecvInterface(ok bool, truthy, falsey <-chan interface{}) <-chan interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvString provides the ternary operator for <-chan string types
func ChanRecvString(ok bool, truthy, falsey <-chan string) <-chan string {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUint provides the ternary operator for <-chan uint types
func ChanRecvUint(ok bool, truthy, falsey <-chan uint) <-chan uint {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUint16 provides the ternary operator for <-chan uint16 types
func ChanRecvUint16(ok bool, truthy, falsey <-chan uint16) <-chan uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUint32 provides the ternary operator for <-chan uint32 types
func ChanRecvUint32(ok bool, truthy, falsey <-chan uint32) <-chan uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUint64 provides the ternary operator for <-chan uint64 types
func ChanRecvUint64(ok bool, truthy, falsey <-chan uint64) <-chan uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUint8 provides the ternary operator for <-chan uint8 types
func ChanRecvUint8(ok bool, truthy, falsey <-chan uint8) <-chan uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUintptr provides the ternary operator for <-chan uintptr types
func ChanRecvUintptr(ok bool, truthy, falsey <-chan uintptr) <-chan uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// ChanRecvUnsafePointer provides the ternary operator for <-chan unsafe.Pointer types
func ChanRecvUnsafePointer(ok bool, truthy, falsey <-chan unsafe.Pointer) <-chan unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendBool provides the ternary operator for chan<- bool types
func ChanSendBool(ok bool, truthy, falsey chan<- bool) chan<- bool {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendComplex128 provides the ternary operator for chan<- complex128 types
func ChanSendComplex128(ok bool, truthy, falsey chan<- complex128) chan<- complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendComplex64 provides the ternary operator for chan<- complex64 types
func ChanSendComplex64(ok bool, truthy, falsey chan<- complex64) chan<- complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendFloat32 provides the ternary operator for chan<- float32 types
func ChanSendFloat32(ok bool, truthy, falsey chan<- float32) chan<- float32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendFloat64 provides the ternary operator for chan<- float64 types
func ChanSendFloat64(ok bool, truthy, falsey chan<- float64) chan<- float64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendInt provides the ternary operator for chan<- int types
func ChanSendInt(ok bool, truthy, falsey chan<- int) chan<- int {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendInt16 provides the ternary operator for chan<- int16 types
func ChanSendInt16(ok bool, truthy, falsey chan<- int16) chan<- int16 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendInt32 provides the ternary operator for chan<- int32 types
func ChanSendInt32(ok bool, truthy, falsey chan<- int32) chan<- int32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendInt64 provides the ternary operator for chan<- int64 types
func ChanSendInt64(ok bool, truthy, falsey chan<- int64) chan<- int64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendInt8 provides the ternary operator for chan<- int8 types
func ChanSendInt8(ok bool, truthy, falsey chan<- int8) chan<- int8 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendInterface provides the ternary operator for chan<- interface{} types
func ChanSendInterface(ok bool, truthy, falsey chan<- interface{}) chan<- interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendString provides the ternary operator for chan<- string types
func ChanSendString(ok bool, truthy, falsey chan<- string) chan<- string {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUint provides the ternary operator for chan<- uint types
func ChanSendUint(ok bool, truthy, falsey chan<- uint) chan<- uint {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUint16 provides the ternary operator for chan<- uint16 types
func ChanSendUint16(ok bool, truthy, falsey chan<- uint16) chan<- uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUint32 provides the ternary operator for chan<- uint32 types
func ChanSendUint32(ok bool, truthy, falsey chan<- uint32) chan<- uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUint64 provides the ternary operator for chan<- uint64 types
func ChanSendUint64(ok bool, truthy, falsey chan<- uint64) chan<- uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUint8 provides the ternary operator for chan<- uint8 types
func ChanSendUint8(ok bool, truthy, falsey chan<- uint8) chan<- uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUintptr provides the ternary operator for chan<- uintptr types
func ChanSendUintptr(ok bool, truthy, falsey chan<- uintptr) chan<- uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// ChanSendUnsafePointer provides the ternary operator for chan<- unsafe.Pointer types
func ChanSendUnsafePointer(ok bool, truthy, falsey chan<- unsafe.Pointer) chan<- unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// ChanString provides the ternary operator for chan string types
func ChanString(ok bool, truthy, falsey chan string) chan string {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUint provides the ternary operator for chan uint types
func ChanUint(ok bool, truthy, falsey chan uint) chan uint {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUint16 provides the ternary operator for chan uint16 types
func ChanUint16(ok bool, truthy, falsey chan uint16) chan uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUint32 provides the ternary operator for chan uint32 types
func ChanUint32(ok bool, truthy, falsey chan uint32) chan uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUint64 provides the ternary operator for chan uint64 types
func ChanUint64(ok bool, truthy, falsey chan uint64) chan uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUint8 provides the ternary operator for chan uint8 types
func ChanUint8(ok bool, truthy, falsey chan uint8) chan uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUintptr provides the ternary operator for chan uintptr types
func ChanUintptr(ok bool, truthy, falsey chan uintptr) chan uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// ChanUnsafePointer provides the ternary operator for chan unsafe.Pointer types
func ChanUnsafePointer(ok bool, truthy, falsey chan unsafe.Pointer) chan unsafe.Pointer {
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

// MapBoolToBool provides the ternary operator for map[bool]bool types
func MapBoolToBool(ok bool, truthy, falsey map[bool]bool) map[bool]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToComplex128 provides the ternary operator for map[bool]complex128 types
func MapBoolToComplex128(ok bool, truthy, falsey map[bool]complex128) map[bool]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToComplex64 provides the ternary operator for map[bool]complex64 types
func MapBoolToComplex64(ok bool, truthy, falsey map[bool]complex64) map[bool]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToFloat32 provides the ternary operator for map[bool]float32 types
func MapBoolToFloat32(ok bool, truthy, falsey map[bool]float32) map[bool]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToFloat64 provides the ternary operator for map[bool]float64 types
func MapBoolToFloat64(ok bool, truthy, falsey map[bool]float64) map[bool]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToInt provides the ternary operator for map[bool]int types
func MapBoolToInt(ok bool, truthy, falsey map[bool]int) map[bool]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToInt16 provides the ternary operator for map[bool]int16 types
func MapBoolToInt16(ok bool, truthy, falsey map[bool]int16) map[bool]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToInt32 provides the ternary operator for map[bool]int32 types
func MapBoolToInt32(ok bool, truthy, falsey map[bool]int32) map[bool]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToInt64 provides the ternary operator for map[bool]int64 types
func MapBoolToInt64(ok bool, truthy, falsey map[bool]int64) map[bool]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToInt8 provides the ternary operator for map[bool]int8 types
func MapBoolToInt8(ok bool, truthy, falsey map[bool]int8) map[bool]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToInterface provides the ternary operator for map[bool]interface{} types
func MapBoolToInterface(ok bool, truthy, falsey map[bool]interface{}) map[bool]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToString provides the ternary operator for map[bool]string types
func MapBoolToString(ok bool, truthy, falsey map[bool]string) map[bool]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUint provides the ternary operator for map[bool]uint types
func MapBoolToUint(ok bool, truthy, falsey map[bool]uint) map[bool]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUint16 provides the ternary operator for map[bool]uint16 types
func MapBoolToUint16(ok bool, truthy, falsey map[bool]uint16) map[bool]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUint32 provides the ternary operator for map[bool]uint32 types
func MapBoolToUint32(ok bool, truthy, falsey map[bool]uint32) map[bool]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUint64 provides the ternary operator for map[bool]uint64 types
func MapBoolToUint64(ok bool, truthy, falsey map[bool]uint64) map[bool]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUint8 provides the ternary operator for map[bool]uint8 types
func MapBoolToUint8(ok bool, truthy, falsey map[bool]uint8) map[bool]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUintptr provides the ternary operator for map[bool]uintptr types
func MapBoolToUintptr(ok bool, truthy, falsey map[bool]uintptr) map[bool]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapBoolToUnsafePointer provides the ternary operator for map[bool]unsafe.Pointer types
func MapBoolToUnsafePointer(ok bool, truthy, falsey map[bool]unsafe.Pointer) map[bool]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToBool provides the ternary operator for map[complex128]bool types
func MapComplex128ToBool(ok bool, truthy, falsey map[complex128]bool) map[complex128]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToComplex128 provides the ternary operator for map[complex128]complex128 types
func MapComplex128ToComplex128(ok bool, truthy, falsey map[complex128]complex128) map[complex128]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToComplex64 provides the ternary operator for map[complex128]complex64 types
func MapComplex128ToComplex64(ok bool, truthy, falsey map[complex128]complex64) map[complex128]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToFloat32 provides the ternary operator for map[complex128]float32 types
func MapComplex128ToFloat32(ok bool, truthy, falsey map[complex128]float32) map[complex128]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToFloat64 provides the ternary operator for map[complex128]float64 types
func MapComplex128ToFloat64(ok bool, truthy, falsey map[complex128]float64) map[complex128]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToInt provides the ternary operator for map[complex128]int types
func MapComplex128ToInt(ok bool, truthy, falsey map[complex128]int) map[complex128]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToInt16 provides the ternary operator for map[complex128]int16 types
func MapComplex128ToInt16(ok bool, truthy, falsey map[complex128]int16) map[complex128]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToInt32 provides the ternary operator for map[complex128]int32 types
func MapComplex128ToInt32(ok bool, truthy, falsey map[complex128]int32) map[complex128]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToInt64 provides the ternary operator for map[complex128]int64 types
func MapComplex128ToInt64(ok bool, truthy, falsey map[complex128]int64) map[complex128]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToInt8 provides the ternary operator for map[complex128]int8 types
func MapComplex128ToInt8(ok bool, truthy, falsey map[complex128]int8) map[complex128]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToInterface provides the ternary operator for map[complex128]interface{} types
func MapComplex128ToInterface(ok bool, truthy, falsey map[complex128]interface{}) map[complex128]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToString provides the ternary operator for map[complex128]string types
func MapComplex128ToString(ok bool, truthy, falsey map[complex128]string) map[complex128]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUint provides the ternary operator for map[complex128]uint types
func MapComplex128ToUint(ok bool, truthy, falsey map[complex128]uint) map[complex128]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUint16 provides the ternary operator for map[complex128]uint16 types
func MapComplex128ToUint16(ok bool, truthy, falsey map[complex128]uint16) map[complex128]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUint32 provides the ternary operator for map[complex128]uint32 types
func MapComplex128ToUint32(ok bool, truthy, falsey map[complex128]uint32) map[complex128]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUint64 provides the ternary operator for map[complex128]uint64 types
func MapComplex128ToUint64(ok bool, truthy, falsey map[complex128]uint64) map[complex128]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUint8 provides the ternary operator for map[complex128]uint8 types
func MapComplex128ToUint8(ok bool, truthy, falsey map[complex128]uint8) map[complex128]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUintptr provides the ternary operator for map[complex128]uintptr types
func MapComplex128ToUintptr(ok bool, truthy, falsey map[complex128]uintptr) map[complex128]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex128ToUnsafePointer provides the ternary operator for map[complex128]unsafe.Pointer types
func MapComplex128ToUnsafePointer(ok bool, truthy, falsey map[complex128]unsafe.Pointer) map[complex128]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToBool provides the ternary operator for map[complex64]bool types
func MapComplex64ToBool(ok bool, truthy, falsey map[complex64]bool) map[complex64]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToComplex128 provides the ternary operator for map[complex64]complex128 types
func MapComplex64ToComplex128(ok bool, truthy, falsey map[complex64]complex128) map[complex64]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToComplex64 provides the ternary operator for map[complex64]complex64 types
func MapComplex64ToComplex64(ok bool, truthy, falsey map[complex64]complex64) map[complex64]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToFloat32 provides the ternary operator for map[complex64]float32 types
func MapComplex64ToFloat32(ok bool, truthy, falsey map[complex64]float32) map[complex64]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToFloat64 provides the ternary operator for map[complex64]float64 types
func MapComplex64ToFloat64(ok bool, truthy, falsey map[complex64]float64) map[complex64]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToInt provides the ternary operator for map[complex64]int types
func MapComplex64ToInt(ok bool, truthy, falsey map[complex64]int) map[complex64]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToInt16 provides the ternary operator for map[complex64]int16 types
func MapComplex64ToInt16(ok bool, truthy, falsey map[complex64]int16) map[complex64]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToInt32 provides the ternary operator for map[complex64]int32 types
func MapComplex64ToInt32(ok bool, truthy, falsey map[complex64]int32) map[complex64]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToInt64 provides the ternary operator for map[complex64]int64 types
func MapComplex64ToInt64(ok bool, truthy, falsey map[complex64]int64) map[complex64]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToInt8 provides the ternary operator for map[complex64]int8 types
func MapComplex64ToInt8(ok bool, truthy, falsey map[complex64]int8) map[complex64]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToInterface provides the ternary operator for map[complex64]interface{} types
func MapComplex64ToInterface(ok bool, truthy, falsey map[complex64]interface{}) map[complex64]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToString provides the ternary operator for map[complex64]string types
func MapComplex64ToString(ok bool, truthy, falsey map[complex64]string) map[complex64]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUint provides the ternary operator for map[complex64]uint types
func MapComplex64ToUint(ok bool, truthy, falsey map[complex64]uint) map[complex64]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUint16 provides the ternary operator for map[complex64]uint16 types
func MapComplex64ToUint16(ok bool, truthy, falsey map[complex64]uint16) map[complex64]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUint32 provides the ternary operator for map[complex64]uint32 types
func MapComplex64ToUint32(ok bool, truthy, falsey map[complex64]uint32) map[complex64]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUint64 provides the ternary operator for map[complex64]uint64 types
func MapComplex64ToUint64(ok bool, truthy, falsey map[complex64]uint64) map[complex64]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUint8 provides the ternary operator for map[complex64]uint8 types
func MapComplex64ToUint8(ok bool, truthy, falsey map[complex64]uint8) map[complex64]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUintptr provides the ternary operator for map[complex64]uintptr types
func MapComplex64ToUintptr(ok bool, truthy, falsey map[complex64]uintptr) map[complex64]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapComplex64ToUnsafePointer provides the ternary operator for map[complex64]unsafe.Pointer types
func MapComplex64ToUnsafePointer(ok bool, truthy, falsey map[complex64]unsafe.Pointer) map[complex64]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToBool provides the ternary operator for map[float32]bool types
func MapFloat32ToBool(ok bool, truthy, falsey map[float32]bool) map[float32]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToComplex128 provides the ternary operator for map[float32]complex128 types
func MapFloat32ToComplex128(ok bool, truthy, falsey map[float32]complex128) map[float32]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToComplex64 provides the ternary operator for map[float32]complex64 types
func MapFloat32ToComplex64(ok bool, truthy, falsey map[float32]complex64) map[float32]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToFloat32 provides the ternary operator for map[float32]float32 types
func MapFloat32ToFloat32(ok bool, truthy, falsey map[float32]float32) map[float32]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToFloat64 provides the ternary operator for map[float32]float64 types
func MapFloat32ToFloat64(ok bool, truthy, falsey map[float32]float64) map[float32]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToInt provides the ternary operator for map[float32]int types
func MapFloat32ToInt(ok bool, truthy, falsey map[float32]int) map[float32]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToInt16 provides the ternary operator for map[float32]int16 types
func MapFloat32ToInt16(ok bool, truthy, falsey map[float32]int16) map[float32]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToInt32 provides the ternary operator for map[float32]int32 types
func MapFloat32ToInt32(ok bool, truthy, falsey map[float32]int32) map[float32]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToInt64 provides the ternary operator for map[float32]int64 types
func MapFloat32ToInt64(ok bool, truthy, falsey map[float32]int64) map[float32]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToInt8 provides the ternary operator for map[float32]int8 types
func MapFloat32ToInt8(ok bool, truthy, falsey map[float32]int8) map[float32]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToInterface provides the ternary operator for map[float32]interface{} types
func MapFloat32ToInterface(ok bool, truthy, falsey map[float32]interface{}) map[float32]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToString provides the ternary operator for map[float32]string types
func MapFloat32ToString(ok bool, truthy, falsey map[float32]string) map[float32]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUint provides the ternary operator for map[float32]uint types
func MapFloat32ToUint(ok bool, truthy, falsey map[float32]uint) map[float32]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUint16 provides the ternary operator for map[float32]uint16 types
func MapFloat32ToUint16(ok bool, truthy, falsey map[float32]uint16) map[float32]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUint32 provides the ternary operator for map[float32]uint32 types
func MapFloat32ToUint32(ok bool, truthy, falsey map[float32]uint32) map[float32]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUint64 provides the ternary operator for map[float32]uint64 types
func MapFloat32ToUint64(ok bool, truthy, falsey map[float32]uint64) map[float32]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUint8 provides the ternary operator for map[float32]uint8 types
func MapFloat32ToUint8(ok bool, truthy, falsey map[float32]uint8) map[float32]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUintptr provides the ternary operator for map[float32]uintptr types
func MapFloat32ToUintptr(ok bool, truthy, falsey map[float32]uintptr) map[float32]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat32ToUnsafePointer provides the ternary operator for map[float32]unsafe.Pointer types
func MapFloat32ToUnsafePointer(ok bool, truthy, falsey map[float32]unsafe.Pointer) map[float32]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToBool provides the ternary operator for map[float64]bool types
func MapFloat64ToBool(ok bool, truthy, falsey map[float64]bool) map[float64]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToComplex128 provides the ternary operator for map[float64]complex128 types
func MapFloat64ToComplex128(ok bool, truthy, falsey map[float64]complex128) map[float64]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToComplex64 provides the ternary operator for map[float64]complex64 types
func MapFloat64ToComplex64(ok bool, truthy, falsey map[float64]complex64) map[float64]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToFloat32 provides the ternary operator for map[float64]float32 types
func MapFloat64ToFloat32(ok bool, truthy, falsey map[float64]float32) map[float64]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToFloat64 provides the ternary operator for map[float64]float64 types
func MapFloat64ToFloat64(ok bool, truthy, falsey map[float64]float64) map[float64]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToInt provides the ternary operator for map[float64]int types
func MapFloat64ToInt(ok bool, truthy, falsey map[float64]int) map[float64]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToInt16 provides the ternary operator for map[float64]int16 types
func MapFloat64ToInt16(ok bool, truthy, falsey map[float64]int16) map[float64]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToInt32 provides the ternary operator for map[float64]int32 types
func MapFloat64ToInt32(ok bool, truthy, falsey map[float64]int32) map[float64]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToInt64 provides the ternary operator for map[float64]int64 types
func MapFloat64ToInt64(ok bool, truthy, falsey map[float64]int64) map[float64]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToInt8 provides the ternary operator for map[float64]int8 types
func MapFloat64ToInt8(ok bool, truthy, falsey map[float64]int8) map[float64]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToInterface provides the ternary operator for map[float64]interface{} types
func MapFloat64ToInterface(ok bool, truthy, falsey map[float64]interface{}) map[float64]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToString provides the ternary operator for map[float64]string types
func MapFloat64ToString(ok bool, truthy, falsey map[float64]string) map[float64]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUint provides the ternary operator for map[float64]uint types
func MapFloat64ToUint(ok bool, truthy, falsey map[float64]uint) map[float64]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUint16 provides the ternary operator for map[float64]uint16 types
func MapFloat64ToUint16(ok bool, truthy, falsey map[float64]uint16) map[float64]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUint32 provides the ternary operator for map[float64]uint32 types
func MapFloat64ToUint32(ok bool, truthy, falsey map[float64]uint32) map[float64]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUint64 provides the ternary operator for map[float64]uint64 types
func MapFloat64ToUint64(ok bool, truthy, falsey map[float64]uint64) map[float64]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUint8 provides the ternary operator for map[float64]uint8 types
func MapFloat64ToUint8(ok bool, truthy, falsey map[float64]uint8) map[float64]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUintptr provides the ternary operator for map[float64]uintptr types
func MapFloat64ToUintptr(ok bool, truthy, falsey map[float64]uintptr) map[float64]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapFloat64ToUnsafePointer provides the ternary operator for map[float64]unsafe.Pointer types
func MapFloat64ToUnsafePointer(ok bool, truthy, falsey map[float64]unsafe.Pointer) map[float64]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToBool provides the ternary operator for map[int16]bool types
func MapInt16ToBool(ok bool, truthy, falsey map[int16]bool) map[int16]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToComplex128 provides the ternary operator for map[int16]complex128 types
func MapInt16ToComplex128(ok bool, truthy, falsey map[int16]complex128) map[int16]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToComplex64 provides the ternary operator for map[int16]complex64 types
func MapInt16ToComplex64(ok bool, truthy, falsey map[int16]complex64) map[int16]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToFloat32 provides the ternary operator for map[int16]float32 types
func MapInt16ToFloat32(ok bool, truthy, falsey map[int16]float32) map[int16]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToFloat64 provides the ternary operator for map[int16]float64 types
func MapInt16ToFloat64(ok bool, truthy, falsey map[int16]float64) map[int16]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToInt provides the ternary operator for map[int16]int types
func MapInt16ToInt(ok bool, truthy, falsey map[int16]int) map[int16]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToInt16 provides the ternary operator for map[int16]int16 types
func MapInt16ToInt16(ok bool, truthy, falsey map[int16]int16) map[int16]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToInt32 provides the ternary operator for map[int16]int32 types
func MapInt16ToInt32(ok bool, truthy, falsey map[int16]int32) map[int16]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToInt64 provides the ternary operator for map[int16]int64 types
func MapInt16ToInt64(ok bool, truthy, falsey map[int16]int64) map[int16]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToInt8 provides the ternary operator for map[int16]int8 types
func MapInt16ToInt8(ok bool, truthy, falsey map[int16]int8) map[int16]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToInterface provides the ternary operator for map[int16]interface{} types
func MapInt16ToInterface(ok bool, truthy, falsey map[int16]interface{}) map[int16]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToString provides the ternary operator for map[int16]string types
func MapInt16ToString(ok bool, truthy, falsey map[int16]string) map[int16]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUint provides the ternary operator for map[int16]uint types
func MapInt16ToUint(ok bool, truthy, falsey map[int16]uint) map[int16]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUint16 provides the ternary operator for map[int16]uint16 types
func MapInt16ToUint16(ok bool, truthy, falsey map[int16]uint16) map[int16]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUint32 provides the ternary operator for map[int16]uint32 types
func MapInt16ToUint32(ok bool, truthy, falsey map[int16]uint32) map[int16]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUint64 provides the ternary operator for map[int16]uint64 types
func MapInt16ToUint64(ok bool, truthy, falsey map[int16]uint64) map[int16]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUint8 provides the ternary operator for map[int16]uint8 types
func MapInt16ToUint8(ok bool, truthy, falsey map[int16]uint8) map[int16]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUintptr provides the ternary operator for map[int16]uintptr types
func MapInt16ToUintptr(ok bool, truthy, falsey map[int16]uintptr) map[int16]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt16ToUnsafePointer provides the ternary operator for map[int16]unsafe.Pointer types
func MapInt16ToUnsafePointer(ok bool, truthy, falsey map[int16]unsafe.Pointer) map[int16]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToBool provides the ternary operator for map[int32]bool types
func MapInt32ToBool(ok bool, truthy, falsey map[int32]bool) map[int32]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToComplex128 provides the ternary operator for map[int32]complex128 types
func MapInt32ToComplex128(ok bool, truthy, falsey map[int32]complex128) map[int32]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToComplex64 provides the ternary operator for map[int32]complex64 types
func MapInt32ToComplex64(ok bool, truthy, falsey map[int32]complex64) map[int32]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToFloat32 provides the ternary operator for map[int32]float32 types
func MapInt32ToFloat32(ok bool, truthy, falsey map[int32]float32) map[int32]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToFloat64 provides the ternary operator for map[int32]float64 types
func MapInt32ToFloat64(ok bool, truthy, falsey map[int32]float64) map[int32]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToInt provides the ternary operator for map[int32]int types
func MapInt32ToInt(ok bool, truthy, falsey map[int32]int) map[int32]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToInt16 provides the ternary operator for map[int32]int16 types
func MapInt32ToInt16(ok bool, truthy, falsey map[int32]int16) map[int32]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToInt32 provides the ternary operator for map[int32]int32 types
func MapInt32ToInt32(ok bool, truthy, falsey map[int32]int32) map[int32]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToInt64 provides the ternary operator for map[int32]int64 types
func MapInt32ToInt64(ok bool, truthy, falsey map[int32]int64) map[int32]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToInt8 provides the ternary operator for map[int32]int8 types
func MapInt32ToInt8(ok bool, truthy, falsey map[int32]int8) map[int32]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToInterface provides the ternary operator for map[int32]interface{} types
func MapInt32ToInterface(ok bool, truthy, falsey map[int32]interface{}) map[int32]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToString provides the ternary operator for map[int32]string types
func MapInt32ToString(ok bool, truthy, falsey map[int32]string) map[int32]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUint provides the ternary operator for map[int32]uint types
func MapInt32ToUint(ok bool, truthy, falsey map[int32]uint) map[int32]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUint16 provides the ternary operator for map[int32]uint16 types
func MapInt32ToUint16(ok bool, truthy, falsey map[int32]uint16) map[int32]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUint32 provides the ternary operator for map[int32]uint32 types
func MapInt32ToUint32(ok bool, truthy, falsey map[int32]uint32) map[int32]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUint64 provides the ternary operator for map[int32]uint64 types
func MapInt32ToUint64(ok bool, truthy, falsey map[int32]uint64) map[int32]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUint8 provides the ternary operator for map[int32]uint8 types
func MapInt32ToUint8(ok bool, truthy, falsey map[int32]uint8) map[int32]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUintptr provides the ternary operator for map[int32]uintptr types
func MapInt32ToUintptr(ok bool, truthy, falsey map[int32]uintptr) map[int32]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt32ToUnsafePointer provides the ternary operator for map[int32]unsafe.Pointer types
func MapInt32ToUnsafePointer(ok bool, truthy, falsey map[int32]unsafe.Pointer) map[int32]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToBool provides the ternary operator for map[int64]bool types
func MapInt64ToBool(ok bool, truthy, falsey map[int64]bool) map[int64]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToComplex128 provides the ternary operator for map[int64]complex128 types
func MapInt64ToComplex128(ok bool, truthy, falsey map[int64]complex128) map[int64]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToComplex64 provides the ternary operator for map[int64]complex64 types
func MapInt64ToComplex64(ok bool, truthy, falsey map[int64]complex64) map[int64]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToFloat32 provides the ternary operator for map[int64]float32 types
func MapInt64ToFloat32(ok bool, truthy, falsey map[int64]float32) map[int64]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToFloat64 provides the ternary operator for map[int64]float64 types
func MapInt64ToFloat64(ok bool, truthy, falsey map[int64]float64) map[int64]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToInt provides the ternary operator for map[int64]int types
func MapInt64ToInt(ok bool, truthy, falsey map[int64]int) map[int64]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToInt16 provides the ternary operator for map[int64]int16 types
func MapInt64ToInt16(ok bool, truthy, falsey map[int64]int16) map[int64]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToInt32 provides the ternary operator for map[int64]int32 types
func MapInt64ToInt32(ok bool, truthy, falsey map[int64]int32) map[int64]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToInt64 provides the ternary operator for map[int64]int64 types
func MapInt64ToInt64(ok bool, truthy, falsey map[int64]int64) map[int64]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToInt8 provides the ternary operator for map[int64]int8 types
func MapInt64ToInt8(ok bool, truthy, falsey map[int64]int8) map[int64]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToInterface provides the ternary operator for map[int64]interface{} types
func MapInt64ToInterface(ok bool, truthy, falsey map[int64]interface{}) map[int64]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToString provides the ternary operator for map[int64]string types
func MapInt64ToString(ok bool, truthy, falsey map[int64]string) map[int64]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUint provides the ternary operator for map[int64]uint types
func MapInt64ToUint(ok bool, truthy, falsey map[int64]uint) map[int64]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUint16 provides the ternary operator for map[int64]uint16 types
func MapInt64ToUint16(ok bool, truthy, falsey map[int64]uint16) map[int64]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUint32 provides the ternary operator for map[int64]uint32 types
func MapInt64ToUint32(ok bool, truthy, falsey map[int64]uint32) map[int64]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUint64 provides the ternary operator for map[int64]uint64 types
func MapInt64ToUint64(ok bool, truthy, falsey map[int64]uint64) map[int64]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUint8 provides the ternary operator for map[int64]uint8 types
func MapInt64ToUint8(ok bool, truthy, falsey map[int64]uint8) map[int64]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUintptr provides the ternary operator for map[int64]uintptr types
func MapInt64ToUintptr(ok bool, truthy, falsey map[int64]uintptr) map[int64]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt64ToUnsafePointer provides the ternary operator for map[int64]unsafe.Pointer types
func MapInt64ToUnsafePointer(ok bool, truthy, falsey map[int64]unsafe.Pointer) map[int64]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToBool provides the ternary operator for map[int8]bool types
func MapInt8ToBool(ok bool, truthy, falsey map[int8]bool) map[int8]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToComplex128 provides the ternary operator for map[int8]complex128 types
func MapInt8ToComplex128(ok bool, truthy, falsey map[int8]complex128) map[int8]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToComplex64 provides the ternary operator for map[int8]complex64 types
func MapInt8ToComplex64(ok bool, truthy, falsey map[int8]complex64) map[int8]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToFloat32 provides the ternary operator for map[int8]float32 types
func MapInt8ToFloat32(ok bool, truthy, falsey map[int8]float32) map[int8]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToFloat64 provides the ternary operator for map[int8]float64 types
func MapInt8ToFloat64(ok bool, truthy, falsey map[int8]float64) map[int8]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToInt provides the ternary operator for map[int8]int types
func MapInt8ToInt(ok bool, truthy, falsey map[int8]int) map[int8]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToInt16 provides the ternary operator for map[int8]int16 types
func MapInt8ToInt16(ok bool, truthy, falsey map[int8]int16) map[int8]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToInt32 provides the ternary operator for map[int8]int32 types
func MapInt8ToInt32(ok bool, truthy, falsey map[int8]int32) map[int8]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToInt64 provides the ternary operator for map[int8]int64 types
func MapInt8ToInt64(ok bool, truthy, falsey map[int8]int64) map[int8]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToInt8 provides the ternary operator for map[int8]int8 types
func MapInt8ToInt8(ok bool, truthy, falsey map[int8]int8) map[int8]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToInterface provides the ternary operator for map[int8]interface{} types
func MapInt8ToInterface(ok bool, truthy, falsey map[int8]interface{}) map[int8]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToString provides the ternary operator for map[int8]string types
func MapInt8ToString(ok bool, truthy, falsey map[int8]string) map[int8]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUint provides the ternary operator for map[int8]uint types
func MapInt8ToUint(ok bool, truthy, falsey map[int8]uint) map[int8]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUint16 provides the ternary operator for map[int8]uint16 types
func MapInt8ToUint16(ok bool, truthy, falsey map[int8]uint16) map[int8]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUint32 provides the ternary operator for map[int8]uint32 types
func MapInt8ToUint32(ok bool, truthy, falsey map[int8]uint32) map[int8]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUint64 provides the ternary operator for map[int8]uint64 types
func MapInt8ToUint64(ok bool, truthy, falsey map[int8]uint64) map[int8]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUint8 provides the ternary operator for map[int8]uint8 types
func MapInt8ToUint8(ok bool, truthy, falsey map[int8]uint8) map[int8]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUintptr provides the ternary operator for map[int8]uintptr types
func MapInt8ToUintptr(ok bool, truthy, falsey map[int8]uintptr) map[int8]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapInt8ToUnsafePointer provides the ternary operator for map[int8]unsafe.Pointer types
func MapInt8ToUnsafePointer(ok bool, truthy, falsey map[int8]unsafe.Pointer) map[int8]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToBool provides the ternary operator for map[int]bool types
func MapIntToBool(ok bool, truthy, falsey map[int]bool) map[int]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToComplex128 provides the ternary operator for map[int]complex128 types
func MapIntToComplex128(ok bool, truthy, falsey map[int]complex128) map[int]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToComplex64 provides the ternary operator for map[int]complex64 types
func MapIntToComplex64(ok bool, truthy, falsey map[int]complex64) map[int]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToFloat32 provides the ternary operator for map[int]float32 types
func MapIntToFloat32(ok bool, truthy, falsey map[int]float32) map[int]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToFloat64 provides the ternary operator for map[int]float64 types
func MapIntToFloat64(ok bool, truthy, falsey map[int]float64) map[int]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToInt provides the ternary operator for map[int]int types
func MapIntToInt(ok bool, truthy, falsey map[int]int) map[int]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToInt16 provides the ternary operator for map[int]int16 types
func MapIntToInt16(ok bool, truthy, falsey map[int]int16) map[int]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToInt32 provides the ternary operator for map[int]int32 types
func MapIntToInt32(ok bool, truthy, falsey map[int]int32) map[int]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToInt64 provides the ternary operator for map[int]int64 types
func MapIntToInt64(ok bool, truthy, falsey map[int]int64) map[int]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToInt8 provides the ternary operator for map[int]int8 types
func MapIntToInt8(ok bool, truthy, falsey map[int]int8) map[int]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToInterface provides the ternary operator for map[int]interface{} types
func MapIntToInterface(ok bool, truthy, falsey map[int]interface{}) map[int]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToString provides the ternary operator for map[int]string types
func MapIntToString(ok bool, truthy, falsey map[int]string) map[int]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUint provides the ternary operator for map[int]uint types
func MapIntToUint(ok bool, truthy, falsey map[int]uint) map[int]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUint16 provides the ternary operator for map[int]uint16 types
func MapIntToUint16(ok bool, truthy, falsey map[int]uint16) map[int]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUint32 provides the ternary operator for map[int]uint32 types
func MapIntToUint32(ok bool, truthy, falsey map[int]uint32) map[int]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUint64 provides the ternary operator for map[int]uint64 types
func MapIntToUint64(ok bool, truthy, falsey map[int]uint64) map[int]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUint8 provides the ternary operator for map[int]uint8 types
func MapIntToUint8(ok bool, truthy, falsey map[int]uint8) map[int]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUintptr provides the ternary operator for map[int]uintptr types
func MapIntToUintptr(ok bool, truthy, falsey map[int]uintptr) map[int]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapIntToUnsafePointer provides the ternary operator for map[int]unsafe.Pointer types
func MapIntToUnsafePointer(ok bool, truthy, falsey map[int]unsafe.Pointer) map[int]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToBool provides the ternary operator for map[interface{}]bool types
func MapInterfaceToBool(ok bool, truthy, falsey map[interface{}]bool) map[interface{}]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToComplex128 provides the ternary operator for map[interface{}]complex128 types
func MapInterfaceToComplex128(ok bool, truthy, falsey map[interface{}]complex128) map[interface{}]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToComplex64 provides the ternary operator for map[interface{}]complex64 types
func MapInterfaceToComplex64(ok bool, truthy, falsey map[interface{}]complex64) map[interface{}]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToFloat32 provides the ternary operator for map[interface{}]float32 types
func MapInterfaceToFloat32(ok bool, truthy, falsey map[interface{}]float32) map[interface{}]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToFloat64 provides the ternary operator for map[interface{}]float64 types
func MapInterfaceToFloat64(ok bool, truthy, falsey map[interface{}]float64) map[interface{}]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToInt provides the ternary operator for map[interface{}]int types
func MapInterfaceToInt(ok bool, truthy, falsey map[interface{}]int) map[interface{}]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToInt16 provides the ternary operator for map[interface{}]int16 types
func MapInterfaceToInt16(ok bool, truthy, falsey map[interface{}]int16) map[interface{}]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToInt32 provides the ternary operator for map[interface{}]int32 types
func MapInterfaceToInt32(ok bool, truthy, falsey map[interface{}]int32) map[interface{}]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToInt64 provides the ternary operator for map[interface{}]int64 types
func MapInterfaceToInt64(ok bool, truthy, falsey map[interface{}]int64) map[interface{}]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToInt8 provides the ternary operator for map[interface{}]int8 types
func MapInterfaceToInt8(ok bool, truthy, falsey map[interface{}]int8) map[interface{}]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToInterface provides the ternary operator for map[interface{}]interface{} types
func MapInterfaceToInterface(ok bool, truthy, falsey map[interface{}]interface{}) map[interface{}]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToString provides the ternary operator for map[interface{}]string types
func MapInterfaceToString(ok bool, truthy, falsey map[interface{}]string) map[interface{}]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUint provides the ternary operator for map[interface{}]uint types
func MapInterfaceToUint(ok bool, truthy, falsey map[interface{}]uint) map[interface{}]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUint16 provides the ternary operator for map[interface{}]uint16 types
func MapInterfaceToUint16(ok bool, truthy, falsey map[interface{}]uint16) map[interface{}]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUint32 provides the ternary operator for map[interface{}]uint32 types
func MapInterfaceToUint32(ok bool, truthy, falsey map[interface{}]uint32) map[interface{}]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUint64 provides the ternary operator for map[interface{}]uint64 types
func MapInterfaceToUint64(ok bool, truthy, falsey map[interface{}]uint64) map[interface{}]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUint8 provides the ternary operator for map[interface{}]uint8 types
func MapInterfaceToUint8(ok bool, truthy, falsey map[interface{}]uint8) map[interface{}]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUintptr provides the ternary operator for map[interface{}]uintptr types
func MapInterfaceToUintptr(ok bool, truthy, falsey map[interface{}]uintptr) map[interface{}]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapInterfaceToUnsafePointer provides the ternary operator for map[interface{}]unsafe.Pointer types
func MapInterfaceToUnsafePointer(ok bool, truthy, falsey map[interface{}]unsafe.Pointer) map[interface{}]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToBool provides the ternary operator for map[string]bool types
func MapStringToBool(ok bool, truthy, falsey map[string]bool) map[string]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToComplex128 provides the ternary operator for map[string]complex128 types
func MapStringToComplex128(ok bool, truthy, falsey map[string]complex128) map[string]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToComplex64 provides the ternary operator for map[string]complex64 types
func MapStringToComplex64(ok bool, truthy, falsey map[string]complex64) map[string]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToFloat32 provides the ternary operator for map[string]float32 types
func MapStringToFloat32(ok bool, truthy, falsey map[string]float32) map[string]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToFloat64 provides the ternary operator for map[string]float64 types
func MapStringToFloat64(ok bool, truthy, falsey map[string]float64) map[string]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToInt provides the ternary operator for map[string]int types
func MapStringToInt(ok bool, truthy, falsey map[string]int) map[string]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToInt16 provides the ternary operator for map[string]int16 types
func MapStringToInt16(ok bool, truthy, falsey map[string]int16) map[string]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToInt32 provides the ternary operator for map[string]int32 types
func MapStringToInt32(ok bool, truthy, falsey map[string]int32) map[string]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToInt64 provides the ternary operator for map[string]int64 types
func MapStringToInt64(ok bool, truthy, falsey map[string]int64) map[string]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToInt8 provides the ternary operator for map[string]int8 types
func MapStringToInt8(ok bool, truthy, falsey map[string]int8) map[string]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToInterface provides the ternary operator for map[string]interface{} types
func MapStringToInterface(ok bool, truthy, falsey map[string]interface{}) map[string]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToString provides the ternary operator for map[string]string types
func MapStringToString(ok bool, truthy, falsey map[string]string) map[string]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUint provides the ternary operator for map[string]uint types
func MapStringToUint(ok bool, truthy, falsey map[string]uint) map[string]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUint16 provides the ternary operator for map[string]uint16 types
func MapStringToUint16(ok bool, truthy, falsey map[string]uint16) map[string]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUint32 provides the ternary operator for map[string]uint32 types
func MapStringToUint32(ok bool, truthy, falsey map[string]uint32) map[string]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUint64 provides the ternary operator for map[string]uint64 types
func MapStringToUint64(ok bool, truthy, falsey map[string]uint64) map[string]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUint8 provides the ternary operator for map[string]uint8 types
func MapStringToUint8(ok bool, truthy, falsey map[string]uint8) map[string]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUintptr provides the ternary operator for map[string]uintptr types
func MapStringToUintptr(ok bool, truthy, falsey map[string]uintptr) map[string]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapStringToUnsafePointer provides the ternary operator for map[string]unsafe.Pointer types
func MapStringToUnsafePointer(ok bool, truthy, falsey map[string]unsafe.Pointer) map[string]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToBool provides the ternary operator for map[uint16]bool types
func MapUint16ToBool(ok bool, truthy, falsey map[uint16]bool) map[uint16]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToComplex128 provides the ternary operator for map[uint16]complex128 types
func MapUint16ToComplex128(ok bool, truthy, falsey map[uint16]complex128) map[uint16]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToComplex64 provides the ternary operator for map[uint16]complex64 types
func MapUint16ToComplex64(ok bool, truthy, falsey map[uint16]complex64) map[uint16]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToFloat32 provides the ternary operator for map[uint16]float32 types
func MapUint16ToFloat32(ok bool, truthy, falsey map[uint16]float32) map[uint16]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToFloat64 provides the ternary operator for map[uint16]float64 types
func MapUint16ToFloat64(ok bool, truthy, falsey map[uint16]float64) map[uint16]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToInt provides the ternary operator for map[uint16]int types
func MapUint16ToInt(ok bool, truthy, falsey map[uint16]int) map[uint16]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToInt16 provides the ternary operator for map[uint16]int16 types
func MapUint16ToInt16(ok bool, truthy, falsey map[uint16]int16) map[uint16]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToInt32 provides the ternary operator for map[uint16]int32 types
func MapUint16ToInt32(ok bool, truthy, falsey map[uint16]int32) map[uint16]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToInt64 provides the ternary operator for map[uint16]int64 types
func MapUint16ToInt64(ok bool, truthy, falsey map[uint16]int64) map[uint16]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToInt8 provides the ternary operator for map[uint16]int8 types
func MapUint16ToInt8(ok bool, truthy, falsey map[uint16]int8) map[uint16]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToInterface provides the ternary operator for map[uint16]interface{} types
func MapUint16ToInterface(ok bool, truthy, falsey map[uint16]interface{}) map[uint16]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToString provides the ternary operator for map[uint16]string types
func MapUint16ToString(ok bool, truthy, falsey map[uint16]string) map[uint16]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUint provides the ternary operator for map[uint16]uint types
func MapUint16ToUint(ok bool, truthy, falsey map[uint16]uint) map[uint16]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUint16 provides the ternary operator for map[uint16]uint16 types
func MapUint16ToUint16(ok bool, truthy, falsey map[uint16]uint16) map[uint16]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUint32 provides the ternary operator for map[uint16]uint32 types
func MapUint16ToUint32(ok bool, truthy, falsey map[uint16]uint32) map[uint16]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUint64 provides the ternary operator for map[uint16]uint64 types
func MapUint16ToUint64(ok bool, truthy, falsey map[uint16]uint64) map[uint16]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUint8 provides the ternary operator for map[uint16]uint8 types
func MapUint16ToUint8(ok bool, truthy, falsey map[uint16]uint8) map[uint16]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUintptr provides the ternary operator for map[uint16]uintptr types
func MapUint16ToUintptr(ok bool, truthy, falsey map[uint16]uintptr) map[uint16]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint16ToUnsafePointer provides the ternary operator for map[uint16]unsafe.Pointer types
func MapUint16ToUnsafePointer(ok bool, truthy, falsey map[uint16]unsafe.Pointer) map[uint16]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToBool provides the ternary operator for map[uint32]bool types
func MapUint32ToBool(ok bool, truthy, falsey map[uint32]bool) map[uint32]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToComplex128 provides the ternary operator for map[uint32]complex128 types
func MapUint32ToComplex128(ok bool, truthy, falsey map[uint32]complex128) map[uint32]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToComplex64 provides the ternary operator for map[uint32]complex64 types
func MapUint32ToComplex64(ok bool, truthy, falsey map[uint32]complex64) map[uint32]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToFloat32 provides the ternary operator for map[uint32]float32 types
func MapUint32ToFloat32(ok bool, truthy, falsey map[uint32]float32) map[uint32]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToFloat64 provides the ternary operator for map[uint32]float64 types
func MapUint32ToFloat64(ok bool, truthy, falsey map[uint32]float64) map[uint32]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToInt provides the ternary operator for map[uint32]int types
func MapUint32ToInt(ok bool, truthy, falsey map[uint32]int) map[uint32]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToInt16 provides the ternary operator for map[uint32]int16 types
func MapUint32ToInt16(ok bool, truthy, falsey map[uint32]int16) map[uint32]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToInt32 provides the ternary operator for map[uint32]int32 types
func MapUint32ToInt32(ok bool, truthy, falsey map[uint32]int32) map[uint32]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToInt64 provides the ternary operator for map[uint32]int64 types
func MapUint32ToInt64(ok bool, truthy, falsey map[uint32]int64) map[uint32]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToInt8 provides the ternary operator for map[uint32]int8 types
func MapUint32ToInt8(ok bool, truthy, falsey map[uint32]int8) map[uint32]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToInterface provides the ternary operator for map[uint32]interface{} types
func MapUint32ToInterface(ok bool, truthy, falsey map[uint32]interface{}) map[uint32]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToString provides the ternary operator for map[uint32]string types
func MapUint32ToString(ok bool, truthy, falsey map[uint32]string) map[uint32]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUint provides the ternary operator for map[uint32]uint types
func MapUint32ToUint(ok bool, truthy, falsey map[uint32]uint) map[uint32]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUint16 provides the ternary operator for map[uint32]uint16 types
func MapUint32ToUint16(ok bool, truthy, falsey map[uint32]uint16) map[uint32]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUint32 provides the ternary operator for map[uint32]uint32 types
func MapUint32ToUint32(ok bool, truthy, falsey map[uint32]uint32) map[uint32]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUint64 provides the ternary operator for map[uint32]uint64 types
func MapUint32ToUint64(ok bool, truthy, falsey map[uint32]uint64) map[uint32]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUint8 provides the ternary operator for map[uint32]uint8 types
func MapUint32ToUint8(ok bool, truthy, falsey map[uint32]uint8) map[uint32]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUintptr provides the ternary operator for map[uint32]uintptr types
func MapUint32ToUintptr(ok bool, truthy, falsey map[uint32]uintptr) map[uint32]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint32ToUnsafePointer provides the ternary operator for map[uint32]unsafe.Pointer types
func MapUint32ToUnsafePointer(ok bool, truthy, falsey map[uint32]unsafe.Pointer) map[uint32]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToBool provides the ternary operator for map[uint64]bool types
func MapUint64ToBool(ok bool, truthy, falsey map[uint64]bool) map[uint64]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToComplex128 provides the ternary operator for map[uint64]complex128 types
func MapUint64ToComplex128(ok bool, truthy, falsey map[uint64]complex128) map[uint64]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToComplex64 provides the ternary operator for map[uint64]complex64 types
func MapUint64ToComplex64(ok bool, truthy, falsey map[uint64]complex64) map[uint64]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToFloat32 provides the ternary operator for map[uint64]float32 types
func MapUint64ToFloat32(ok bool, truthy, falsey map[uint64]float32) map[uint64]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToFloat64 provides the ternary operator for map[uint64]float64 types
func MapUint64ToFloat64(ok bool, truthy, falsey map[uint64]float64) map[uint64]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToInt provides the ternary operator for map[uint64]int types
func MapUint64ToInt(ok bool, truthy, falsey map[uint64]int) map[uint64]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToInt16 provides the ternary operator for map[uint64]int16 types
func MapUint64ToInt16(ok bool, truthy, falsey map[uint64]int16) map[uint64]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToInt32 provides the ternary operator for map[uint64]int32 types
func MapUint64ToInt32(ok bool, truthy, falsey map[uint64]int32) map[uint64]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToInt64 provides the ternary operator for map[uint64]int64 types
func MapUint64ToInt64(ok bool, truthy, falsey map[uint64]int64) map[uint64]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToInt8 provides the ternary operator for map[uint64]int8 types
func MapUint64ToInt8(ok bool, truthy, falsey map[uint64]int8) map[uint64]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToInterface provides the ternary operator for map[uint64]interface{} types
func MapUint64ToInterface(ok bool, truthy, falsey map[uint64]interface{}) map[uint64]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToString provides the ternary operator for map[uint64]string types
func MapUint64ToString(ok bool, truthy, falsey map[uint64]string) map[uint64]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUint provides the ternary operator for map[uint64]uint types
func MapUint64ToUint(ok bool, truthy, falsey map[uint64]uint) map[uint64]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUint16 provides the ternary operator for map[uint64]uint16 types
func MapUint64ToUint16(ok bool, truthy, falsey map[uint64]uint16) map[uint64]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUint32 provides the ternary operator for map[uint64]uint32 types
func MapUint64ToUint32(ok bool, truthy, falsey map[uint64]uint32) map[uint64]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUint64 provides the ternary operator for map[uint64]uint64 types
func MapUint64ToUint64(ok bool, truthy, falsey map[uint64]uint64) map[uint64]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUint8 provides the ternary operator for map[uint64]uint8 types
func MapUint64ToUint8(ok bool, truthy, falsey map[uint64]uint8) map[uint64]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUintptr provides the ternary operator for map[uint64]uintptr types
func MapUint64ToUintptr(ok bool, truthy, falsey map[uint64]uintptr) map[uint64]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint64ToUnsafePointer provides the ternary operator for map[uint64]unsafe.Pointer types
func MapUint64ToUnsafePointer(ok bool, truthy, falsey map[uint64]unsafe.Pointer) map[uint64]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToBool provides the ternary operator for map[uint8]bool types
func MapUint8ToBool(ok bool, truthy, falsey map[uint8]bool) map[uint8]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToComplex128 provides the ternary operator for map[uint8]complex128 types
func MapUint8ToComplex128(ok bool, truthy, falsey map[uint8]complex128) map[uint8]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToComplex64 provides the ternary operator for map[uint8]complex64 types
func MapUint8ToComplex64(ok bool, truthy, falsey map[uint8]complex64) map[uint8]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToFloat32 provides the ternary operator for map[uint8]float32 types
func MapUint8ToFloat32(ok bool, truthy, falsey map[uint8]float32) map[uint8]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToFloat64 provides the ternary operator for map[uint8]float64 types
func MapUint8ToFloat64(ok bool, truthy, falsey map[uint8]float64) map[uint8]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToInt provides the ternary operator for map[uint8]int types
func MapUint8ToInt(ok bool, truthy, falsey map[uint8]int) map[uint8]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToInt16 provides the ternary operator for map[uint8]int16 types
func MapUint8ToInt16(ok bool, truthy, falsey map[uint8]int16) map[uint8]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToInt32 provides the ternary operator for map[uint8]int32 types
func MapUint8ToInt32(ok bool, truthy, falsey map[uint8]int32) map[uint8]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToInt64 provides the ternary operator for map[uint8]int64 types
func MapUint8ToInt64(ok bool, truthy, falsey map[uint8]int64) map[uint8]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToInt8 provides the ternary operator for map[uint8]int8 types
func MapUint8ToInt8(ok bool, truthy, falsey map[uint8]int8) map[uint8]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToInterface provides the ternary operator for map[uint8]interface{} types
func MapUint8ToInterface(ok bool, truthy, falsey map[uint8]interface{}) map[uint8]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToString provides the ternary operator for map[uint8]string types
func MapUint8ToString(ok bool, truthy, falsey map[uint8]string) map[uint8]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUint provides the ternary operator for map[uint8]uint types
func MapUint8ToUint(ok bool, truthy, falsey map[uint8]uint) map[uint8]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUint16 provides the ternary operator for map[uint8]uint16 types
func MapUint8ToUint16(ok bool, truthy, falsey map[uint8]uint16) map[uint8]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUint32 provides the ternary operator for map[uint8]uint32 types
func MapUint8ToUint32(ok bool, truthy, falsey map[uint8]uint32) map[uint8]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUint64 provides the ternary operator for map[uint8]uint64 types
func MapUint8ToUint64(ok bool, truthy, falsey map[uint8]uint64) map[uint8]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUint8 provides the ternary operator for map[uint8]uint8 types
func MapUint8ToUint8(ok bool, truthy, falsey map[uint8]uint8) map[uint8]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUintptr provides the ternary operator for map[uint8]uintptr types
func MapUint8ToUintptr(ok bool, truthy, falsey map[uint8]uintptr) map[uint8]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUint8ToUnsafePointer provides the ternary operator for map[uint8]unsafe.Pointer types
func MapUint8ToUnsafePointer(ok bool, truthy, falsey map[uint8]unsafe.Pointer) map[uint8]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToBool provides the ternary operator for map[uint]bool types
func MapUintToBool(ok bool, truthy, falsey map[uint]bool) map[uint]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToComplex128 provides the ternary operator for map[uint]complex128 types
func MapUintToComplex128(ok bool, truthy, falsey map[uint]complex128) map[uint]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToComplex64 provides the ternary operator for map[uint]complex64 types
func MapUintToComplex64(ok bool, truthy, falsey map[uint]complex64) map[uint]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToFloat32 provides the ternary operator for map[uint]float32 types
func MapUintToFloat32(ok bool, truthy, falsey map[uint]float32) map[uint]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToFloat64 provides the ternary operator for map[uint]float64 types
func MapUintToFloat64(ok bool, truthy, falsey map[uint]float64) map[uint]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToInt provides the ternary operator for map[uint]int types
func MapUintToInt(ok bool, truthy, falsey map[uint]int) map[uint]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToInt16 provides the ternary operator for map[uint]int16 types
func MapUintToInt16(ok bool, truthy, falsey map[uint]int16) map[uint]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToInt32 provides the ternary operator for map[uint]int32 types
func MapUintToInt32(ok bool, truthy, falsey map[uint]int32) map[uint]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToInt64 provides the ternary operator for map[uint]int64 types
func MapUintToInt64(ok bool, truthy, falsey map[uint]int64) map[uint]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToInt8 provides the ternary operator for map[uint]int8 types
func MapUintToInt8(ok bool, truthy, falsey map[uint]int8) map[uint]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToInterface provides the ternary operator for map[uint]interface{} types
func MapUintToInterface(ok bool, truthy, falsey map[uint]interface{}) map[uint]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToString provides the ternary operator for map[uint]string types
func MapUintToString(ok bool, truthy, falsey map[uint]string) map[uint]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUint provides the ternary operator for map[uint]uint types
func MapUintToUint(ok bool, truthy, falsey map[uint]uint) map[uint]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUint16 provides the ternary operator for map[uint]uint16 types
func MapUintToUint16(ok bool, truthy, falsey map[uint]uint16) map[uint]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUint32 provides the ternary operator for map[uint]uint32 types
func MapUintToUint32(ok bool, truthy, falsey map[uint]uint32) map[uint]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUint64 provides the ternary operator for map[uint]uint64 types
func MapUintToUint64(ok bool, truthy, falsey map[uint]uint64) map[uint]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUint8 provides the ternary operator for map[uint]uint8 types
func MapUintToUint8(ok bool, truthy, falsey map[uint]uint8) map[uint]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUintptr provides the ternary operator for map[uint]uintptr types
func MapUintToUintptr(ok bool, truthy, falsey map[uint]uintptr) map[uint]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintToUnsafePointer provides the ternary operator for map[uint]unsafe.Pointer types
func MapUintToUnsafePointer(ok bool, truthy, falsey map[uint]unsafe.Pointer) map[uint]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToBool provides the ternary operator for map[uintptr]bool types
func MapUintptrToBool(ok bool, truthy, falsey map[uintptr]bool) map[uintptr]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToComplex128 provides the ternary operator for map[uintptr]complex128 types
func MapUintptrToComplex128(ok bool, truthy, falsey map[uintptr]complex128) map[uintptr]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToComplex64 provides the ternary operator for map[uintptr]complex64 types
func MapUintptrToComplex64(ok bool, truthy, falsey map[uintptr]complex64) map[uintptr]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToFloat32 provides the ternary operator for map[uintptr]float32 types
func MapUintptrToFloat32(ok bool, truthy, falsey map[uintptr]float32) map[uintptr]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToFloat64 provides the ternary operator for map[uintptr]float64 types
func MapUintptrToFloat64(ok bool, truthy, falsey map[uintptr]float64) map[uintptr]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToInt provides the ternary operator for map[uintptr]int types
func MapUintptrToInt(ok bool, truthy, falsey map[uintptr]int) map[uintptr]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToInt16 provides the ternary operator for map[uintptr]int16 types
func MapUintptrToInt16(ok bool, truthy, falsey map[uintptr]int16) map[uintptr]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToInt32 provides the ternary operator for map[uintptr]int32 types
func MapUintptrToInt32(ok bool, truthy, falsey map[uintptr]int32) map[uintptr]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToInt64 provides the ternary operator for map[uintptr]int64 types
func MapUintptrToInt64(ok bool, truthy, falsey map[uintptr]int64) map[uintptr]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToInt8 provides the ternary operator for map[uintptr]int8 types
func MapUintptrToInt8(ok bool, truthy, falsey map[uintptr]int8) map[uintptr]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToInterface provides the ternary operator for map[uintptr]interface{} types
func MapUintptrToInterface(ok bool, truthy, falsey map[uintptr]interface{}) map[uintptr]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToString provides the ternary operator for map[uintptr]string types
func MapUintptrToString(ok bool, truthy, falsey map[uintptr]string) map[uintptr]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUint provides the ternary operator for map[uintptr]uint types
func MapUintptrToUint(ok bool, truthy, falsey map[uintptr]uint) map[uintptr]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUint16 provides the ternary operator for map[uintptr]uint16 types
func MapUintptrToUint16(ok bool, truthy, falsey map[uintptr]uint16) map[uintptr]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUint32 provides the ternary operator for map[uintptr]uint32 types
func MapUintptrToUint32(ok bool, truthy, falsey map[uintptr]uint32) map[uintptr]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUint64 provides the ternary operator for map[uintptr]uint64 types
func MapUintptrToUint64(ok bool, truthy, falsey map[uintptr]uint64) map[uintptr]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUint8 provides the ternary operator for map[uintptr]uint8 types
func MapUintptrToUint8(ok bool, truthy, falsey map[uintptr]uint8) map[uintptr]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUintptr provides the ternary operator for map[uintptr]uintptr types
func MapUintptrToUintptr(ok bool, truthy, falsey map[uintptr]uintptr) map[uintptr]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUintptrToUnsafePointer provides the ternary operator for map[uintptr]unsafe.Pointer types
func MapUintptrToUnsafePointer(ok bool, truthy, falsey map[uintptr]unsafe.Pointer) map[uintptr]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToBool provides the ternary operator for map[unsafe.Pointer]bool types
func MapUnsafePointerToBool(ok bool, truthy, falsey map[unsafe.Pointer]bool) map[unsafe.Pointer]bool {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToComplex128 provides the ternary operator for map[unsafe.Pointer]complex128 types
func MapUnsafePointerToComplex128(ok bool, truthy, falsey map[unsafe.Pointer]complex128) map[unsafe.Pointer]complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToComplex64 provides the ternary operator for map[unsafe.Pointer]complex64 types
func MapUnsafePointerToComplex64(ok bool, truthy, falsey map[unsafe.Pointer]complex64) map[unsafe.Pointer]complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToFloat32 provides the ternary operator for map[unsafe.Pointer]float32 types
func MapUnsafePointerToFloat32(ok bool, truthy, falsey map[unsafe.Pointer]float32) map[unsafe.Pointer]float32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToFloat64 provides the ternary operator for map[unsafe.Pointer]float64 types
func MapUnsafePointerToFloat64(ok bool, truthy, falsey map[unsafe.Pointer]float64) map[unsafe.Pointer]float64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToInt provides the ternary operator for map[unsafe.Pointer]int types
func MapUnsafePointerToInt(ok bool, truthy, falsey map[unsafe.Pointer]int) map[unsafe.Pointer]int {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToInt16 provides the ternary operator for map[unsafe.Pointer]int16 types
func MapUnsafePointerToInt16(ok bool, truthy, falsey map[unsafe.Pointer]int16) map[unsafe.Pointer]int16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToInt32 provides the ternary operator for map[unsafe.Pointer]int32 types
func MapUnsafePointerToInt32(ok bool, truthy, falsey map[unsafe.Pointer]int32) map[unsafe.Pointer]int32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToInt64 provides the ternary operator for map[unsafe.Pointer]int64 types
func MapUnsafePointerToInt64(ok bool, truthy, falsey map[unsafe.Pointer]int64) map[unsafe.Pointer]int64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToInt8 provides the ternary operator for map[unsafe.Pointer]int8 types
func MapUnsafePointerToInt8(ok bool, truthy, falsey map[unsafe.Pointer]int8) map[unsafe.Pointer]int8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToInterface provides the ternary operator for map[unsafe.Pointer]interface{} types
func MapUnsafePointerToInterface(ok bool, truthy, falsey map[unsafe.Pointer]interface{}) map[unsafe.Pointer]interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToString provides the ternary operator for map[unsafe.Pointer]string types
func MapUnsafePointerToString(ok bool, truthy, falsey map[unsafe.Pointer]string) map[unsafe.Pointer]string {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUint provides the ternary operator for map[unsafe.Pointer]uint types
func MapUnsafePointerToUint(ok bool, truthy, falsey map[unsafe.Pointer]uint) map[unsafe.Pointer]uint {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUint16 provides the ternary operator for map[unsafe.Pointer]uint16 types
func MapUnsafePointerToUint16(ok bool, truthy, falsey map[unsafe.Pointer]uint16) map[unsafe.Pointer]uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUint32 provides the ternary operator for map[unsafe.Pointer]uint32 types
func MapUnsafePointerToUint32(ok bool, truthy, falsey map[unsafe.Pointer]uint32) map[unsafe.Pointer]uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUint64 provides the ternary operator for map[unsafe.Pointer]uint64 types
func MapUnsafePointerToUint64(ok bool, truthy, falsey map[unsafe.Pointer]uint64) map[unsafe.Pointer]uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUint8 provides the ternary operator for map[unsafe.Pointer]uint8 types
func MapUnsafePointerToUint8(ok bool, truthy, falsey map[unsafe.Pointer]uint8) map[unsafe.Pointer]uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUintptr provides the ternary operator for map[unsafe.Pointer]uintptr types
func MapUnsafePointerToUintptr(ok bool, truthy, falsey map[unsafe.Pointer]uintptr) map[unsafe.Pointer]uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// MapUnsafePointerToUnsafePointer provides the ternary operator for map[unsafe.Pointer]unsafe.Pointer types
func MapUnsafePointerToUnsafePointer(ok bool, truthy, falsey map[unsafe.Pointer]unsafe.Pointer) map[unsafe.Pointer]unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// PtrBool provides the ternary operator for *bool types
func PtrBool(ok bool, truthy, falsey *bool) *bool {
	if ok {
		return truthy
	}
	return falsey
}

// PtrComplex128 provides the ternary operator for *complex128 types
func PtrComplex128(ok bool, truthy, falsey *complex128) *complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrComplex64 provides the ternary operator for *complex64 types
func PtrComplex64(ok bool, truthy, falsey *complex64) *complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrFloat32 provides the ternary operator for *float32 types
func PtrFloat32(ok bool, truthy, falsey *float32) *float32 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrFloat64 provides the ternary operator for *float64 types
func PtrFloat64(ok bool, truthy, falsey *float64) *float64 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrInt provides the ternary operator for *int types
func PtrInt(ok bool, truthy, falsey *int) *int {
	if ok {
		return truthy
	}
	return falsey
}

// PtrInt16 provides the ternary operator for *int16 types
func PtrInt16(ok bool, truthy, falsey *int16) *int16 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrInt32 provides the ternary operator for *int32 types
func PtrInt32(ok bool, truthy, falsey *int32) *int32 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrInt64 provides the ternary operator for *int64 types
func PtrInt64(ok bool, truthy, falsey *int64) *int64 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrInt8 provides the ternary operator for *int8 types
func PtrInt8(ok bool, truthy, falsey *int8) *int8 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrInterface provides the ternary operator for *interface{} types
func PtrInterface(ok bool, truthy, falsey *interface{}) *interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// PtrString provides the ternary operator for *string types
func PtrString(ok bool, truthy, falsey *string) *string {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUint provides the ternary operator for *uint types
func PtrUint(ok bool, truthy, falsey *uint) *uint {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUint16 provides the ternary operator for *uint16 types
func PtrUint16(ok bool, truthy, falsey *uint16) *uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUint32 provides the ternary operator for *uint32 types
func PtrUint32(ok bool, truthy, falsey *uint32) *uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUint64 provides the ternary operator for *uint64 types
func PtrUint64(ok bool, truthy, falsey *uint64) *uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUint8 provides the ternary operator for *uint8 types
func PtrUint8(ok bool, truthy, falsey *uint8) *uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUintptr provides the ternary operator for *uintptr types
func PtrUintptr(ok bool, truthy, falsey *uintptr) *uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// PtrUnsafePointer provides the ternary operator for *unsafe.Pointer types
func PtrUnsafePointer(ok bool, truthy, falsey *unsafe.Pointer) *unsafe.Pointer {
	if ok {
		return truthy
	}
	return falsey
}

// SliceBool provides the ternary operator for []bool types
func SliceBool(ok bool, truthy, falsey []bool) []bool {
	if ok {
		return truthy
	}
	return falsey
}

// SliceComplex128 provides the ternary operator for []complex128 types
func SliceComplex128(ok bool, truthy, falsey []complex128) []complex128 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceComplex64 provides the ternary operator for []complex64 types
func SliceComplex64(ok bool, truthy, falsey []complex64) []complex64 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceFloat32 provides the ternary operator for []float32 types
func SliceFloat32(ok bool, truthy, falsey []float32) []float32 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceFloat64 provides the ternary operator for []float64 types
func SliceFloat64(ok bool, truthy, falsey []float64) []float64 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceInt provides the ternary operator for []int types
func SliceInt(ok bool, truthy, falsey []int) []int {
	if ok {
		return truthy
	}
	return falsey
}

// SliceInt16 provides the ternary operator for []int16 types
func SliceInt16(ok bool, truthy, falsey []int16) []int16 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceInt32 provides the ternary operator for []int32 types
func SliceInt32(ok bool, truthy, falsey []int32) []int32 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceInt64 provides the ternary operator for []int64 types
func SliceInt64(ok bool, truthy, falsey []int64) []int64 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceInt8 provides the ternary operator for []int8 types
func SliceInt8(ok bool, truthy, falsey []int8) []int8 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceInterface provides the ternary operator for []interface{} types
func SliceInterface(ok bool, truthy, falsey []interface{}) []interface{} {
	if ok {
		return truthy
	}
	return falsey
}

// SliceString provides the ternary operator for []string types
func SliceString(ok bool, truthy, falsey []string) []string {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUint provides the ternary operator for []uint types
func SliceUint(ok bool, truthy, falsey []uint) []uint {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUint16 provides the ternary operator for []uint16 types
func SliceUint16(ok bool, truthy, falsey []uint16) []uint16 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUint32 provides the ternary operator for []uint32 types
func SliceUint32(ok bool, truthy, falsey []uint32) []uint32 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUint64 provides the ternary operator for []uint64 types
func SliceUint64(ok bool, truthy, falsey []uint64) []uint64 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUint8 provides the ternary operator for []uint8 types
func SliceUint8(ok bool, truthy, falsey []uint8) []uint8 {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUintptr provides the ternary operator for []uintptr types
func SliceUintptr(ok bool, truthy, falsey []uintptr) []uintptr {
	if ok {
		return truthy
	}
	return falsey
}

// SliceUnsafePointer provides the ternary operator for []unsafe.Pointer types
func SliceUnsafePointer(ok bool, truthy, falsey []unsafe.Pointer) []unsafe.Pointer {
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
