package pkg

func SliceToSet[T comparable](s []T) map[T]struct{} {
	return SliceToMap[T, struct{}, T](s, func(t T) (T, struct{}) {
		return t, struct{}{}
	})
}

func SetToSlice[T comparable](m map[T]struct{}) []T {
	return MapToSlice[T, struct{}, T](m, func(t T, s struct{}) T {
		return t
	})
}
