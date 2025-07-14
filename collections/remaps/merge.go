package remaps

func Merge[K comparable, V any](m ...map[K]V) map[K]V {
	m2 := map[K]V{}
	for _, c := range m {
		for k, v := range c {
			m2[k] = v
		}
	}
	return m2
}
