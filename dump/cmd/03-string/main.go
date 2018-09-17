package main

import (
	"encoding/binary"
	"encoding/hex"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func main() {
	println("dumping string:")
	var x string
	x = "Hello, Gophers!"

	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	print("  dump: ", hex.Dump(mem))

	first := mem[:8]  // pointer to data
	second := mem[8:] // length of data

	dataPtr := binary.LittleEndian.Uint64(first)
	length := binary.LittleEndian.Uint64(second)

	println("  data address:", dataPtr)
	println("  length:", length)
	dataMem := rawAccess(unsafe.Pointer(uintptr(dataPtr)), uintptr(length))
	print("  data dump: ", hex.Dump(dataMem))
	println("  string:", string(dataMem))

	mem[8] = 5

	println("  original:", x)
}
