## lessons

### makeing_a_map.go

1) `map[string]string` is a reference type
2) make creates an instance of it
3) address is allocated when make is called.

```golang
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
```
