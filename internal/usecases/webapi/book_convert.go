package webapi

import (
	"slices"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) bookConvert(bookFull entities.BookContainer, attributes map[string]entities.Attribute) entities.BookToWeb {
	book := entities.BookToWeb{
		Book:       bookFull.Book,
		Pages:      make([]entities.PreviewPage, len(bookFull.Pages)),
		Attributes: convertBookAttributes(attributes, bookFull.Attributes),
		Size:       bookFull.Size,
	}

	if len(bookFull.PagesWithHash) > 0 {
		for i, page := range bookFull.PagesWithHash {
			_, hasDeadHash := bookFull.DeadHashOnPage[page.PageNumber]

			book.Pages[i] = page.ToPreview()
			book.Pages[i].HasDeadHash = &hasDeadHash
		}
	} else {
		for i, page := range bookFull.Pages {
			_, hasDeadHash := bookFull.DeadHashOnPage[page.PageNumber]

			book.Pages[i] = page.ToPreview()
			book.Pages[i].HasDeadHash = &hasDeadHash
		}
	}

	if len(book.Pages) > 0 {
		for _, page := range book.Pages {
			if page.PageNumber == entities.PageNumberForPreview {
				book.PreviewPage = page

				break
			}
		}
	}

	for _, attr := range book.Attributes {
		if attr.Code == "tag" { // FIXME: отказаться от такого хардкода
			book.Tags = attr.Values

			break
		}
	}

	slices.SortStableFunc(book.Pages, func(a, b entities.PreviewPage) int {
		return a.PageNumber - b.PageNumber
	})

	return book
}

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
