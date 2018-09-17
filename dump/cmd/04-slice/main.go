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
	println("dumping slice:")
	var x []int32
	x = make([]int32, 0, 8)
	x = append(x, 1, 2, 3, 4)

	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))
	println("  dump:", strings.Replace(strings.Trim(hex.Dump(mem), "\n"), "\n", "\n        ", -1))

	first := mem[:8]    // pointer to data
	second := mem[8:16] // length
	third := mem[16:]   // capacity

	dataPtr := binary.LittleEndian.Uint64(first)
	length := binary.LittleEndian.Uint64(second)
	capacity := binary.LittleEndian.Uint64(third)

	println("  data address:", dataPtr)
	println("  length:", length)
	println("  capacity:", capacity)
	dataMem := rawAccess(unsafe.Pointer(uintptr(dataPtr)), uintptr(length)*unsafe.Sizeof(x[0]))
	print("  data dump: ", hex.Dump(dataMem))

	y := [4]int32{5, 6, 7, 8}

	binary.LittleEndian.PutUint64(first, uint64(uintptr(unsafe.Pointer(&y))))

	fmt.Printf("  original slice: %v\n", x)
}
