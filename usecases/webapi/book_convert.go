package webapi

import (
	"slices"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func convertBookAttributes(attributes map[string]core.Attribute, bookAttributes map[string][]string) []bff.AttributeToWeb {
	result := pkg.MapToSlice(bookAttributes, func(code string, values []string) bff.AttributeToWeb {
		return bff.AttributeToWeb{
			Code:   code,
			Name:   attributes[code].PluralName,
			Values: values,
		}
	})

	slices.SortStableFunc(result, func(a, b bff.AttributeToWeb) int {
		return attributes[a.Code].Order - attributes[b.Code].Order
	})

	return result
}

func convertAttributes(attributes []core.Attribute) map[string]core.Attribute {
	return pkg.SliceToMap(attributes, func(attribute core.Attribute) (string, core.Attribute) {
		return attribute.Code, attribute
	})
}
