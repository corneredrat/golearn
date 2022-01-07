package main

import (
	"fmt"
	"unsafe"
)

type sliceHeader struct {
	pointer  uintptr
	len      uint
	capacity uint
}

func main() {
	var nilSlice []int
	emptySlice := []int{}

	fmt.Printf("nil slice: %v\nempty slice: %v\n", nilSlice, emptySlice)
	fmt.Printf("nil slice ptr:%p\nempty slice ptr: %p\n", nilSlice, emptySlice)
	nilSliceHeader := (*sliceHeader)(unsafe.Pointer(&nilSlice))
	fmt.Printf("nil slice's header: %v\n", nilSliceHeader)

	emptySliceHeader := (*sliceHeader)(unsafe.Pointer(&emptySlice))
	fmt.Printf("empty slice's header: %v\n", emptySliceHeader)

	intptr := emptySliceHeader.pointer
	intval := (*int)(unsafe.Pointer(intptr))
	fmt.Printf("%v -> %v\n", intval, *intval)

	// fmt.Printf("nil slice val:%v\nempty slice val: %v\n", nilSlice[0], emptySlice[0]) ----> panic: runtime error: index out of range [0] with length 0
	// https://go.dev/play/p/jm2p6L6WCG ---> reflecting on slice data type
}
