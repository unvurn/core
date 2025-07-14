package remaps

func EachFunc[K comparable, V, T any](s map[K]V, f func(K, V) T) map[K]T {
	s2 := make(map[K]T, len(s))
	for k, v := range s {
		s2[k] = f(k, v)
	}
	return s2
}

func EachValueFunc[K comparable, V, T any](s map[K]V, f func(K, V) T) []T {
	s2 := make([]T, 0, len(s))
	for k, v := range s {
		s2 = append(s2, f(k, v))
	}
	return s2
}
