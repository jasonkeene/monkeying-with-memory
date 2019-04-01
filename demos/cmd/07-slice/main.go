package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
}

func main() {
	x := make([]int32, 0, 8)
	x = append(x, 1, 2, 3, 4)

	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	println("dumping slice:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", fmt.Sprint(x))
	println("  dump:", strings.Replace(strings.Trim(hex.Dump(mem), "\n"), "\n", "\n        ", -1))

	first := mem[:8]    // pointer to data
	second := mem[8:16] // length
	third := mem[16:]   // capacity

	dataPtr := uintptr(binary.LittleEndian.Uint64(first))
	length := int(binary.LittleEndian.Uint64(second))
	capacity := int(binary.LittleEndian.Uint64(third))

	dataMem := rawAccess(unsafe.Pointer(dataPtr), uintptr(length)*unsafe.Sizeof(x[0]))

	println("dumping slice data:")
	println("  type:", fmt.Sprintf("%T", dataMem))
	println("  address:", unsafe.Pointer(dataPtr))
	println("  length:", length)
	println("  capacity:", capacity)
	print("  dump: ", hex.Dump(dataMem))

	y := [4]int32{5, 6, 7, 8}

	binary.LittleEndian.PutUint64(first, uint64(uintptr(unsafe.Pointer(&y))))

	println("dumping slice with mutation:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", fmt.Sprint(x))
	println("  dump:", strings.Replace(strings.Trim(hex.Dump(mem), "\n"), "\n", "\n        ", -1))
}
