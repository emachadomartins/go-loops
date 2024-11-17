package goloops

type filterFn[T any] func(T, int) bool

func Filter[T any](arr []T, fn filterFn[T]) []T {
	var newArray []T
	for i, v := range arr {
		validate := fn(v, i)
		if validate {
			newArray = append(newArray, v)
		}
	}
	return newArray
}
