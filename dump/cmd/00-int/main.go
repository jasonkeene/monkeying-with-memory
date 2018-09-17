package main

import (
	"encoding/hex"
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
	var x int
	x = 'G'

	// get a slice to that place in memory
	mem := rawAccess(unsafe.Pointer(&x), unsafe.Sizeof(x))

	// dump it
	println("dumping that section of memory")
	println(hex.Dump(mem))

	// mutate it and dump it again
	println("changing the case:")
	mem[0] = mem[0] ^ ' ' // this operation toggles the case of any ascii char
	println(hex.Dump(mem))
}
