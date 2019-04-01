package main

import "sync/atomic"

type A struct {
	A int32 // 4 bytes
	B int64 // 8 bytes
	C int32 // 4 bytes
} // Total: 16 bytes

func main() {
	var a A
	atomic.AddInt64(&a.B, 1)
}
