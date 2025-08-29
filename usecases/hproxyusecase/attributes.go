package hproxyusecase

import (
	"context"
	"fmt"
	"slices"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) handleAttributes(
	ctx context.Context,
	oldAttrs []hproxymodel.BookAttribute,
) ([]hproxymodel.BookAttribute, error) {
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

func (uc *UseCase) setMassloadToAttributes(
	ctx context.Context,
	newAttrs []hproxymodel.BookAttribute,
) error {
	for i, attrs := range newAttrs {
		for j, attr := range attrs.Values {
			mlByAttr, err := uc.storage.MassloadsByAttribute(ctx, attrs.Code, attr.Name)
			if err != nil {
				return fmt.Errorf("storage: get massloads by attribute: %w", err)
			}

			if len(mlByAttr) > 0 {
				newAttrs[i].Values[j].MassloadsByName = pkg.Map(
					mlByAttr,
					func(ml massloadmodel.Massload) hproxymodel.MassloadInfo {
						return hproxymodel.MassloadInfo{
							ID:   ml.ID,
							Name: ml.Name,
						}
					},
				)
			}

			if attr.ExtURL == nil {
				continue
			}

			mlByURL, err := uc.storage.MassloadsByExternalLink(ctx, *attr.ExtURL)
			if err != nil {
				return fmt.Errorf("storage: get massloads by external link: %w", err)
			}

			if len(mlByURL) > 0 {
				newAttrs[i].Values[j].MassloadsByURL = pkg.Map(
					mlByURL,
					func(ml massloadmodel.Massload) hproxymodel.MassloadInfo {
						return hproxymodel.MassloadInfo{
							ID:   ml.ID,
							Name: ml.Name,
						}
					},
				)
			}
		}
	}

	return nil
}

func convertAttributes(attributes []core.Attribute) map[string]core.Attribute {
	return pkg.SliceToMap(attributes, func(attribute core.Attribute) (string, core.Attribute) {
		return attribute.Code, attribute
	})
}
