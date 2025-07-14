package reslices

func EachFunc[T, T1 any](s []T, f func(T) T1) []T1 {
	s2 := make([]T1, 0, len(s))
	for _, e := range s {
		s2 = append(s2, f(e))
	}
	return s2
}

func EachManyFunc[T, T1 any](s []T, f func(T) []T1) []T1 {
	s2 := make([]T1, 0, len(s))
	for _, e := range s {
		s2 = append(s2, f(e)...)
	}
	return s2
}

func EachMapFunc[E any, K comparable](s []E, f func(E) K) map[K]E {
	m := make(map[K]E, len(s))
	for _, e := range s {
		m[f(e)] = e
	}
	return m
}
