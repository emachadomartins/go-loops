# go-loops

Module Library to implement functions like map, forEach, reduce and filter in Go

## Functions

### Map

This function receives as a parameter an Array or a Slice and a function that executes something with the array position as a parameter and return other value. 

This function returns a new Array with the same size than original array but with the type passed as a generic to the function.

```go
package main

import (
	"fmt"
	"github.com/emachadomartins/go-loops"
)

func multiplyTo10(value int, index int) int {
	return value * 10
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	nArr := goloops.Map(arr, multiplyTo10)

	fmt.Println(nArr)
}
```

### ForEach


This function receives as a parameter an Array or a Slice and a function that executes something with the array position as a parameter and returns no value.

This function has no return, is only used for a execution in list.

```go
package main

import (
	"fmt"
	"github.com/emachadomartins/go-loops"
)

func printValues(value int, index int) {
	fmt.Printf("The value of position %v is: %v\n", index, value)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	goloops.ForEach(arr, printValues)
}
```



