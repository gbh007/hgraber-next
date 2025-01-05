package webapi

import (
	"slices"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) bookConvert(bookFull entities.BookFull, attributes map[string]entities.Attribute) entities.BookToWeb {
	book := entities.BookToWeb{
		Book:       bookFull.Book,
		Pages:      make([]entities.PageWithDeadHash, len(bookFull.Pages)),
		Attributes: convertBookAttributes(attributes, bookFull.Attributes),
		Size:       bookFull.Size,
	}

	for i, page := range bookFull.Pages {
		_, hasDeadHash := bookFull.DeadHashOnPage[page.PageNumber]

		book.Pages[i] = entities.PageWithDeadHash{
			Page:        page,
			HasDeadHash: hasDeadHash,
		}
	}

	if len(book.Pages) > 0 {
		book.ParsedPages = true

		for _, page := range book.Pages {
			if page.PageNumber == entities.PageNumberForPreview {
				book.PreviewPage = page.Page

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

	if len(book.Tags) > 8 { // FIXME: отказаться от такой логике на сервере
		book.Tags = book.Tags[:8]
		book.HasMoreTags = true
	}

	slices.SortStableFunc(book.Pages, func(a, b entities.PageWithDeadHash) int {
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
