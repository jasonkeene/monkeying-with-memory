package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func main() {
	x := [4]int32{1, 2, 3, 4}

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping array of int32:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", fmt.Sprint(x))
	print("  dump: ", hex.Dump(mem))
}
