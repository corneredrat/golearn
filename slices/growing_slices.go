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

	a := []int{1, 2, 3, 4, 5}
	b := a[3:4:4]

	aSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Printf("%v, %v\n", b, bSliceHeader)

	b = append(b, 6)
	fmt.Printf("%v, %v\n", a, aSliceHeader)
	fmt.Printf("%v, %v\n", b, bSliceHeader)

}
