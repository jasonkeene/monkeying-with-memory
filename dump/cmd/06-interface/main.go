package main

import (
	"encoding/hex"
	"reflect"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

type someError struct{}

func (*someError) Error() string { return "" }

func typedNilError() *someError { return nil }

func errorValue() *someError { return &someError{} }

func main() {
	println("dumping nil error:")
	var x error

	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	print("  nil dump: ", hex.Dump(mem))
	println("  x == nil:", x == nil)

	println()
	println("dumping typed nil error:")
	x = typedNilError()

	print("  typed nil dump: ", hex.Dump(mem))
	println("  x == nil:", x == nil)
	println("  reflect.ValueOf(x).IsNil():", reflect.ValueOf(x).IsNil())

	println()
	println("dumping error value:")
	x = errorValue()

	print("  value dump: ", hex.Dump(mem))
	println("  x == nil:", x == nil)
	println("  reflect.ValueOf(x).IsNil():", reflect.ValueOf(x).IsNil())
}
