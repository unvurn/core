package reslices

import "slices"

func FilterFunc[T any](s []T, f func(T) bool) []T {
	s2 := make([]T, 0, len(s))
	for _, e := range s {
		if f(e) {
			s2 = append(s2, e)
		}
	}
	return slices.Clip(s2)
}
