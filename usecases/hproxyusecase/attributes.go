package hproxyusecase

import (
	"context"
	"fmt"
	"slices"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) handleAttributes(ctx context.Context, oldAttrs []hproxymodel.BookAttribute) ([]hproxymodel.BookAttribute, error) {
	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: get attributes info: %w", err)
	}

	remaps, err := uc.storage.AttributeRemaps(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: get attributes remaps: %w", err)
	}

	remaper := core.NewAttributeRemaper(remaps, uc.remapToLower)

	attrNames := convertAttributes(attributesInfo)
	newAttrs := make([]hproxymodel.BookAttribute, 0, len(oldAttrs))

	remapped := make(map[string][]hproxymodel.BookAttributeValue, len(oldAttrs))

	for _, oldAttr := range oldAttrs {
		newAttr := hproxymodel.BookAttribute{
			Code: oldAttr.Code,
			Name: attrNames[oldAttr.Code].PluralName,
		}

		for _, oldValue := range oldAttr.Values {
			if !uc.autoRemap {
				newAttr.Values = append(newAttr.Values, oldValue)

				continue
			}

			toCode, toValue, ok := remaper.RemapOne(oldAttr.Code, oldValue.ExtName)
			if !ok {
				newAttr.Values = append(newAttr.Values, oldValue)

				continue
			}

			newValue := hproxymodel.BookAttributeValue{
				ExtName: oldValue.ExtName,
				ExtURL:  oldValue.ExtURL,
				Name:    toValue,
			}

			if newAttr.Code != toCode {
				remapped[toCode] = append(remapped[toCode], newValue)
			} else {
				newAttr.Values = append(newAttr.Values, newValue)
			}
		}

		newAttrs = append(newAttrs, newAttr)
	}

	for i := range newAttrs {
		attr := newAttrs[i]
		attr.Values = append(attr.Values, remapped[attr.Code]...)
		newAttrs[i] = attr
		delete(remapped, attr.Code)
	}

	for code, values := range remapped {
		newAttrs = append(newAttrs, hproxymodel.BookAttribute{
			Code:   code,
			Name:   attrNames[code].PluralName,
			Values: values,
		})
	}

	newAttrs = pkg.SliceFilter(newAttrs, func(ba hproxymodel.BookAttribute) bool {
		return len(ba.Values) > 0
	})

	slices.SortStableFunc(newAttrs, func(a, b hproxymodel.BookAttribute) int {
		return attrNames[a.Code].Order - attrNames[b.Code].Order
	})

	return newAttrs, nil
}

func convertAttributes(attributes []core.Attribute) map[string]core.Attribute {
	return pkg.SliceToMap(attributes, func(attribute core.Attribute) (string, core.Attribute) {
		return attribute.Code, attribute
	})
}
