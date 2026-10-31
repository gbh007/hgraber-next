package pkg

func SliceToMap[K comparable, V, T any](s []T, cv func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(s))

	for _, e := range s {
		k, v := cv(e)
		result[k] = v
	}

	return result
}

func MapToSlice[K comparable, V, T any](m map[K]V, cv func(K, V) T) []T {
	result := make([]T, 0, len(m))

	for k, v := range m {
		e := cv(k, v)
		result = append(result, e)
	}

	return result
}
