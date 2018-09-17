package pkg

import (
	"syscall"
	"unsafe"
)

func AllowWrites(p uintptr) {
	syscall.Mprotect(page(p), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
}

func page(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}
