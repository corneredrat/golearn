package main

import (
	"fmt"
)

func main(){
	var mymap 	= make(map[string]string)
	//var mymap2 	= map[string]string throws error: type map[string]string is not an expression
	fmt.Printf("%v",mymap)
	fmt.Printf("%p",&mymap)
}