package entities

import (
	"hgnext/internal/pkg"
)

type AttributeVariant struct {
	Code  string
	Value string
	Count int
}

func MergeAttributeMap(a, b map[string][]string) map[string][]string {
	result := make(map[string][]string, max(len(a), len(b)))

	for code, values := range a {
		result[code] = values
	}

	for code, values := range b {
		result[code] = pkg.Unique(result[code], values)
	}

	for code, values := range result {
		if len(values) == 0 {
			delete(result, code)
		}
	}

	return result
}
