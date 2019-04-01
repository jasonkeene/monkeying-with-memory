package main

import (
	"encoding/hex"
	"fmt"
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
	x := unaligned{
		a: 1,
		b: 2,
		c: 3,
	}

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping unaligned struct:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", fmt.Sprint(x))
	println("  dump:", strings.Replace(strings.Trim(hex.Dump(mem), "\n"), "\n", "\n        ", -1))

	y := aligned{
		a: 1,
		b: 2,
		c: 3,
	}

	mem = rawAccess(unsafe.Pointer(&y), unsafe.Sizeof(y))

	println("dumping aligned struct:")
	println("  type:", fmt.Sprintf("%T", y))
	println("  address:", unsafe.Pointer(&y))
	println("  sizeof:", unsafe.Sizeof(y))
	println("  value:", fmt.Sprint(y))
	print("  dump: ", hex.Dump(mem))
}
