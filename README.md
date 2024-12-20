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


### Filter


This function receives as a parameter an Array or a Slice and a function that executes something with the array position as a parameter and returns boolean.

This function returns only the values that the condition return the value **true**.

```go
package main

import (
	"fmt"
	"github.com/emachadomartins/go-loops"
)


func checkMultipleOf5(value int, index int) bool {
	condition := value % 5 == 0
	if condition {
		fmt.Printf("%v: %v\n", index, value)
	}
	return condition
}

func main() {
	arr := []int{1, 5, 8, 25, 78, 23, 55, 10, 20}

	filterArr := goloops.Filter(arr, checkMultipleOf5)

	fmt.Println(filterArr)
}

```
### Reduce


This function receives as a parameter an Array or a Slice, a initialValue and a function that executes something with the array position and converts to the type of initialValue

This function transforms the array into a single value.

```go
package main

import (
	"fmt"
	"github.com/emachadomartins/go-loops"
)

type NumberRange struct {
	Max float32
	Min float32
}

func validateRange(previous NumberRange, value float32, index int) NumberRange {
	var maxValue float32
	var minValue float32

	if value > previous.Max {
		fmt.Printf("Value %v in index %v is bigger than %v\n", value, index, previous.Max)
		maxValue = value
	} else {
		maxValue = previous.Max
	}
	if value < previous.Min {
		fmt.Printf("Value %v in index %v is lower than %v\n", value, index, previous.Min)
		minValue = value
	} else {
		minValue = previous.Min
	}

	return NumberRange{
		Max: maxValue,
		Min: minValue,
	}
}

func main() {
	arr := []float32{8.3, 12.4, 0.5, 80.12, -31.2, -90.5, 228.12, 34.9, 100.25}

	var initialValue NumberRange = NumberRange{arr[0], arr[0]}

	finalValue := goloops.Reduce(initialValue, arr, validateRange)

	fmt.Printf("Max value is :%v\n", finalValue.Max)
	fmt.Printf("Min value is :%v\n", finalValue.Min)
}

```



