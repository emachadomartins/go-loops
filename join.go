package goloops

import "fmt"

func Join[T comparable](arr []T, sep string) string {
	s := fmt.Sprint(arr[0])
	for _, v := range arr[1:] {
		s = fmt.Sprintf("%v%v%v", s, sep, v)
	}
	return s
}
