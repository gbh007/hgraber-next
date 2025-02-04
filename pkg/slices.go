package pkg

import (
	"fmt"
)

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

func Unique[V comparable](s ...[]V) []V {
	tmp := make(map[V]struct{}, len(s))

	for _, arr := range s {
		for _, e := range arr {
			tmp[e] = struct{}{}
		}
	}

	result := make([]V, 0, len(tmp))

	for k := range tmp {
		result = append(result, k)
	}

	return result
}

func Batching[T any](a []T, size int) [][]T {
	if len(a) == 0 {
		return nil
	}

	if len(a) <= size {
		return [][]T{a}
	}

	result := make([][]T, 0, len(a)/size+1)

	for len(a) > 0 {
		l := min(size, len(a))

		result = append(result, a[:l])
		a = a[l:]
	}

	return result
}
