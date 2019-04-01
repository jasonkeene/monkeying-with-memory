package main

import "unsafe"

type A struct {
	A int32 // 4 bytes
	B int64 // 8 bytes
	C int32 // 4 bytes
} // Total: 16 bytes

type B struct {
	A int32 // 4 bytes
	C int32 // 4 bytes
	B int64 // 8 bytes
} // Total: 16 bytes

func main() {
	var (
		a [100]A
		b [100]B
	)
	println(unsafe.Sizeof(a))
	println(unsafe.Sizeof(b))
}
