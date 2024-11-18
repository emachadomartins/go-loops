package goloops

type reduceFn[T any, R any] func(R, T, int) R

func Reduce[T any, R any](initialValue R, arr []T, fn reduceFn[T, R]) R {
	var c R = initialValue
	for i, v := range arr {
		c = fn(c, v, i)
	}
	return c
}
