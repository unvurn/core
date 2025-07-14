package remaps

func FilterFunc[K comparable, V any](s map[K]V, f func(K, V) bool) map[K]V {
	s2 := make(map[K]V, len(s))
	for k, v := range s {
		if f(k, v) {
			s2[k] = v
		}
	}
	return s2
}
