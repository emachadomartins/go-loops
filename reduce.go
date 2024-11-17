package main

import "fmt"

type reduceFn[T any, R any] func(R, T, int) R

func Reduce[T any, R any](initialValue R, arr []T, fn reduceFn[T, R]) R {
	var c R = initialValue
	for i, v := range arr {
		c = fn(c, v, i)
	}
	return c
}

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

	finalValue := Reduce(initialValue, arr, validateRange)

	fmt.Printf("Max value is :%v\n", finalValue.Max)
	fmt.Printf("Min value is :%v\n", finalValue.Min)
}
