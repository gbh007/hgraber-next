package webapi

import (
	"slices"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func convertBookAttributes(attributes map[string]entities.Attribute, bookAttributes map[string][]string) []entities.AttributeToWeb {
	result := pkg.MapToSlice(bookAttributes, func(code string, values []string) entities.AttributeToWeb {
		return entities.AttributeToWeb{
			Code:   code,
			Name:   attributes[code].PluralName,
			Values: values,
		}
	})

	slices.SortStableFunc(result, func(a, b entities.AttributeToWeb) int {
		return attributes[a.Code].Order - attributes[b.Code].Order
	})

	return result
}

func convertAttributes(attributes []entities.Attribute) map[string]entities.Attribute {
	return pkg.SliceToMap(attributes, func(attribute entities.Attribute) (string, entities.Attribute) {
		return attribute.Code, attribute
	})
}
