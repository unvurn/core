package reslices

import "slices"

func ContainsAny[T comparable](s, s1 []T) bool {
	return slices.ContainsFunc(s, func(e T) bool {
		return slices.Contains(s1, e)
	})
}

func ContainsAnyFunc[T, U comparable](s []T, s1 []U, f func(T, U) bool) bool {
	return slices.ContainsFunc(s, func(e T) bool {
		return slices.ContainsFunc(s1, func(e1 U) bool {
			return f(e, e1)
		})
	})
}
