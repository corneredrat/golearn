package main

import "fmt"

func main() {
	parentSlice := []string{"bengaluru", "mysuru", "chikkaballapura", "peresandra"}
	childSlice := parentSlice[1:2:3]
	fmt.Printf("values: %v, %v\n", parentSlice, childSlice)
	fmt.Printf("values: %p, %p\n", &parentSlice, childSlice)

}
