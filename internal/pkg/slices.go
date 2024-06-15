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
