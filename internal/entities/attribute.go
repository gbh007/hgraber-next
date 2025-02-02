package entities

import (
	"time"

	"github.com/gbh007/hgraber-next/internal/pkg"
)

const PossibleAttributeCount = 7

type AttributeVariant struct {
	Code  string
	Value string
	Count int
}

type Attribute struct {
	Code        string
	Name        string
	PluralName  string
	Order       int
	Description string
}

type AttributeColor struct {
	Code            string
	Value           string
	TextColor       string
	BackgroundColor string
	CreatedAt       time.Time
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

func AttributesValuesDiff(a, b []string) (aUniq, both, bUniq []string) {
	bMap := make(map[string]struct{})

	for _, value := range b {
		bMap[value] = struct{}{}
	}

	aMap := make(map[string]struct{})

	for _, value := range a {
		aMap[value] = struct{}{}

		_, ok := bMap[value]
		if ok {
			both = append(both, value)
		} else {
			aUniq = append(aUniq, value)
		}
	}

	for _, value := range b {
		_, ok := aMap[value]
		if !ok {
			bUniq = append(bUniq, value)
		}
	}

	return
}
