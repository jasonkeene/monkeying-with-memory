package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func main() {
	x := "Hello, Gophers!"

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping string:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))

	first := mem[:8]  // pointer to data
	second := mem[8:] // length of data

	dataPtr := uintptr(binary.LittleEndian.Uint64(first))
	length := int(binary.LittleEndian.Uint64(second))

	dataMem := rawAccess(unsafe.Pointer(uintptr(dataPtr)), uintptr(length))

	println("dumping string data:")
	println("  type:", fmt.Sprintf("%T", dataMem))
	println("  address:", unsafe.Pointer(dataPtr))
	println("  sizeof:", length)
	println("  value:", string(dataMem))
	print("  dump: ", hex.Dump(dataMem))

	mem[8] = 5 // mutate the length in the string header

	println("dumping string with mutation:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))
}
