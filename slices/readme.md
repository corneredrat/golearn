```golang
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

```


```shell
go run ./growing_slices.go

[4], &{824633844024 1 1}
[1 2 3 4 5], &{824633844000 5 5}
[4 6], &{824633835808 2 2}
```