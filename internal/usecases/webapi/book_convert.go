package webapi

import (
	"slices"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) bookConvert(bookFull entities.BookFull) entities.BookToWeb {
	book := entities.BookToWeb{
		Book:  bookFull.Book,
		Pages: bookFull.Pages,
		Attributes: pkg.MapToSlice(bookFull.Attributes, func(code string, values []string) entities.AttributeToWeb {
			return entities.AttributeToWeb{
				Code:   code,
				Name:   attributeDisplayName(code),
				Values: values,
			}
		}),
		Size: bookFull.Size,
	}

	if len(book.Pages) > 0 {
		book.ParsedPages = true
		book.PreviewPage = book.Pages[0]
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

	slices.SortStableFunc(book.Pages, func(a, b entities.Page) int {
		return a.PageNumber - b.PageNumber
	})

	slices.SortStableFunc(book.Attributes, func(a, b entities.AttributeToWeb) int {
		return attributeOrder(a.Code) - attributeOrder(b.Code)
	})

	return book
}

// FIXME: удалить и брать из базы
func attributeDisplayName(code string) string {
	switch code {
	case "author":
		return "Авторы"
	case "category":
		return "Категории"
	case "character":
		return "Персонажи"
	case "group":
		return "Группы"
	case "language":
		return "Языки"
	case "parody":
		return "Пародии"
	case "tag":
		return "Тэги"
	default:
		return code
	}
}

// FIXME: удалить и брать из базы
func attributeOrder(code string) int {
	switch code {
	case "author":
		return 3
	case "category":
		return 2
	case "character":
		return 4
	case "group":
		return 5
	case "language":
		return 6
	case "parody":
		return 7
	case "tag":
		return 1
	default:
		return 999
	}
}
