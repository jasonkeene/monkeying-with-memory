package main

import (
	"encoding/hex"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func dumpInt() {
	println("dumping int:")
	var x int
	x = -1
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	println("  dump:", hex.Dump(mem))
}

func dumpUint() {
	println("dumping uint:")
	var x uint
	x = 0xfeedface
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	println("  dump:", hex.Dump(mem))
}

func dumpFloat64() {
	println("dumping float64:")
	var x float64
	x = 1.2345
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	println("  dump:", hex.Dump(mem))
}

func dumpComplex128() {
	println("dumping complex128:")
	var x complex128
	x = 1 + 1i
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	println("  dump:", hex.Dump(mem))
}

func main() {
	dumpInt()
	dumpUint()
	dumpFloat64()
	dumpComplex128()
}
