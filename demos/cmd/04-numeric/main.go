package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func dumpInt() {
	x := -1

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping int:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))
}

func dumpUint() {
	var x uint = 0xfeedface

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping uint:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))
}

func dumpFloat64() {
	x := 1.2345

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping float64:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))
}

func dumpComplex128() {
	x := 1 + 1i

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping complex128:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))
}

func main() {
	dumpInt()
	dumpUint()
	dumpFloat64()
	dumpComplex128()
}
