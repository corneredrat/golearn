## Learnings

#### main.go

1) Map is reference type. No need to pass it as a pointer, its always passed as pointer, and any changes in a fn will be reflected in calling fn.

```golang
package main

import (
	"fmt"
)

func main () {
	str := "hi there"
	mymap := map[string]string {"Hi":"There"}
	func (str *string) {
		fmt.Printf("%v %p\n",str,str)
	}(&str)
	func (str *string) {
		fmt.Printf("%v %p\n",*str,str)
	}(&str)
	/*func (mymap2 *map[string]string) {
		fmt.Printf("%v %p\n",mymap2,mymap2)
	}(mymap) 
	//cannot use mymap (type map[string]string) as type *map[string]string in argument to func literal
	*/ 
	func (mymap2 *map[string]string) {
		fmt.Printf("%v %p\n",mymap2,mymap2)
		//*mymap2["Hi"] = "Lol" mymap2["Hi"] (type *map[string]string does not support indexing)
		//mymap2["Hi"] = "Lol" mymap2["Hi"] (type *map[string]string does not support indexing)
	}(&mymap)
	func (mymap2 *map[string]string) {
		fmt.Printf("%v %p\n",mymap2,mymap2)
	}(&mymap)
	func (mymap2 map[string]string) {
		fmt.Printf("%v %p\n",mymap2,mymap2)
		mymap2["Hi"] = "Lol"
		fmt.Printf("%v %p\n",mymap2,mymap2)
	}(mymap)
	func (mymap2 map[string]string) {
		fmt.Printf("%v %p\n",mymap2,mymap2)
	}(mymap)
}
```