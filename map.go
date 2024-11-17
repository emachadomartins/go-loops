package goloops

type mapFn[T any, R any] func(T, int) R

func Map[T any, R any](arr []T, fn mapFn[T, R]) []R {
	newArray := make([]R, len(arr))
	for i, v := range arr {
		newArray = append(newArray, fn(v, i))
	}
	return newArray
}
