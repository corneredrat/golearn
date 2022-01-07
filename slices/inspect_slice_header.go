package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type sliceHeader struct {
	pointer  uintptr
	len      uint
	capacity uint
}

func main() {
	mySlice := []string{"bengaluru", "mysuru", "chikkaballapura", "peresandra"}
	mySliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&mySlice))
	fmt.Printf("Original Silce\nheader: %v, value: %v\n\n", mySliceHeader, mySlice)

	newSlice := mySlice[2:3:3]
	newSliceHeader := (*sliceHeader)(unsafe.Pointer(&newSlice))
	fmt.Printf("New Slice\nheader:%v, value: %v\n\n", newSliceHeader, newSlice)

	myStringPtr := (*string)(unsafe.Pointer(newSliceHeader.pointer))
	fmt.Printf("%p", myStringPtr)
	myString := (string)(*myStringPtr)
	fmt.Printf("\nstr: %v", myString)
}
