package main

import (
	"encoding/hex"
	"strings"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

type unaligned struct {
	a int32
	b int64
	c int32
}

type aligned struct {
	a int32
	c int32
	b int64
}

func main() {
	println("dumping unaligned struct:")
	x := unaligned{a: 1, b: 2, c: 3}

	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	println("  dump:", strings.Replace(strings.Trim(hex.Dump(mem), "\n"), "\n", "\n        ", -1))
	println()

	println("dumping aligned struct:")
	y := aligned{a: 1, b: 2, c: 3}

	println("  address:", unsafe.Pointer(&y))
	println("  sizeof:", unsafe.Sizeof(y))
	mem = rawAccess(unsafe.Pointer(&y), unsafe.Sizeof(y))
	print("  dump: ", hex.Dump(mem))
}
