package main

import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

func rawAccess(p unsafe.Pointer, len uintptr) []byte {
	return (*(*[0xFF]byte)(p))[:len]
	// cast pointer to an array pointer:
	//     (*[0xFF]byte)(p)
	// dereference that pointer
	//     *(*[0xFF]byte)(p)
	// convert the array into a slice
	//     (*(*[0xFF]byte)(p))[:]
	// limit the length of the slice to sizeof
	//     (*(*[0xFF]byte)(p))[:len]
}

func main() {
	// declare and initialize local variable
	var x int = 'G'

	// get a slice to that place in memory
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	// dump it
	println("original int:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))

	// mutate it and dump it again
	mem[0] = mem[0] ^ ' ' // this operation toggles the case of any ascii char
	println("mutated int:")
	println("  type:", fmt.Sprintf("%T", x))
	println("  address:", unsafe.Pointer(&x))
	println("  sizeof:", unsafe.Sizeof(x))
	println("  value:", x)
	print("  dump: ", hex.Dump(mem))
}
