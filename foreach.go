package goloops

type forEachFn[T any] func(T, int)

func ForEach[T any, R any](arr []T, fn forEachFn[T]) {
	for i, v := range arr {
		fn(v, i)
	}
}
