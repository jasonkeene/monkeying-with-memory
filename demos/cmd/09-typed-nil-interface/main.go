package main

import (
	"encoding/hex"
	"fmt"
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
	var x error

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping nil error:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  x == nil:", x == nil)
	print("  dump: ", hex.Dump(mem))

	x = typedNilError()

	println("dumping typed nil error:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  x == nil:", x == nil)
	println("  reflect.ValueOf(x).IsNil():", reflect.ValueOf(x).IsNil())
	print("  dump: ", hex.Dump(mem))

	x = errorValue()

	println("dumping error value:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  x == nil:", x == nil)
	println("  reflect.ValueOf(x).IsNil():", reflect.ValueOf(x).IsNil())
	print("  dump: ", hex.Dump(mem))
}
