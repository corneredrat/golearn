package main

import (
	"fmt"
)

func main() {
	var initializedVar string
	shortDeclarationVar := "This is a string" // variable is set to type of the object returned in RHS implicitly.

	fmt.Println(initializedVar)
	fmt.Println(shortDeclarationVar)

	initializedVar = "This is a new string"
	fmt.Println(initializedVar)

}
