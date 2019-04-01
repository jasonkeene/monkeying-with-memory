package main

import "fmt"

func main() {
	var (
		x      uint32 = 0xFFFFFFFF
		result uint64 = uint64(x) * uint64(x)
	)
	fmt.Printf("% x", result)
}
