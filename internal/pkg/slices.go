package pkg

import "fmt"

func Map[A, B any](a []A, c func(A) B) []B {
	b := make([]B, len(a))
	for i, v := range a {
		b[i] = c(v)
	}

	return b
}

func MapWithError[A, B any](a []A, c func(A) (B, error)) (b []B, err error) {
	b = make([]B, len(a))
	for i, v := range a {
		b[i], err = c(v)
		if err != nil {
			return nil, fmt.Errorf("iter %d: %w", i, err)
		}
	}

	return b, nil
}

func SliceFilter[T any](s []T, f func(T) bool) []T {
	out := make([]T, 0, len(s))

	for _, v := range s {
		if !f(v) {
			continue
		}

		out = append(out, v)
	}

	return out
}

func SliceReduce[T, V any](s []T, f func(sum V, e T) V) V {
	var v V

	for _, e := range s {
		v = f(v, e)
	}

	return v
}
