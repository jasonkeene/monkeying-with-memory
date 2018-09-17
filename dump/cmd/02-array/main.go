package main

import (
	"encoding/hex"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func main() {
	println("dumping array:")
	var x [4]int32
	x = [4]int32{1, 2, 3, 4}

	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	print("  dump: ", hex.Dump(mem))
}
