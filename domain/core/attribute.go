package core

import (
	"strings"
	"time"

	"github.com/gbh007/hgraber-next/pkg"
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

type AttributeRemap struct {
	Code      string
	Value     string
	ToCode    string
	ToValue   string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (ar AttributeRemap) IsDelete() bool {
	return ar.ToCode == "" || ar.ToValue == ""
}

func (ar AttributeRemap) IsNoRemap() bool {
	return ar.Code == ar.ToCode && ar.Value == ar.ToValue
}

type AttributeRemaper struct {
	rules   map[string]map[string]AttributeRemap
	toLower bool
}

func NewAttributeRemaper(remaps []AttributeRemap, toLower bool) AttributeRemaper {
	rmp := AttributeRemaper{
		rules:   make(map[string]map[string]AttributeRemap, PossibleAttributeCount),
		toLower: toLower,
	}

	for _, remap := range remaps {
		if _, ok := rmp.rules[remap.Code]; !ok {
			rmp.rules[remap.Code] = make(map[string]AttributeRemap)
		}

		rmp.rules[remap.Code][remap.Value] = remap
	}

	return rmp
}

func (rmp AttributeRemaper) Remap(origin map[string][]string) map[string][]string {
	attributes := make(map[string][]string, PossibleAttributeCount)

	for code, values := range origin {
		for _, value := range values {
			toCode, toValue, ok := rmp.RemapOne(code, value)
			if !ok {
				continue
			}

			attributes[toCode] = append(attributes[toCode], toValue)
		}
	}

	for code := range attributes {
		attributes[code] = pkg.Unique(attributes[code])
	}

	for code, values := range attributes {
		if len(values) == 0 {
			delete(attributes, code)
		}
	}

	return attributes
}

func (rmp AttributeRemaper) RemapOne(code, value string) (string, string, bool) {
	v := value
	if rmp.toLower {
		v = strings.ToLower(v)
	}

	remap, ok := rmp.rules[code][value]
	if !ok || remap.IsNoRemap() {
		return code, v, true
	}

	if remap.IsDelete() {
		return "", "", false
	}

	toV := remap.ToValue
	if rmp.toLower {
		toV = strings.ToLower(remap.ToValue)
	}

	return remap.ToCode, toV, true
}
