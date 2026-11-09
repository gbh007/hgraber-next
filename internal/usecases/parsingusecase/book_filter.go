package parsingusecase

import (
	"regexp"
	"slices"
	"strings"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

// bookMatchesFilter проверит попадание разобранной книги под фильтр.
// Семантика условий аналогична SQL версии построения запроса (см. build_books_filter.go).
// Учитываются только Fields.Name и Fields.Attributes, остальное игнорируется.
//
//nolint:cyclop // switch по всем типам условий
func bookMatchesFilter(filter core.BookFilter, book hproxymodel.Book) bool {
	if filter.Fields.Name == "" && len(filter.Fields.Attributes) == 0 {
		return true
	}

	if filter.Fields.Name != "" && !matchLike(filter.Fields.Name, book.Name) {
		return false
	}

	for _, attrFilter := range filter.Fields.Attributes {
		values := parsedAttributeValues(book, attrFilter.Code)

		switch attrFilter.Type {
		case core.BookFilterAttributeTypeLike:
			if len(attrFilter.Values) == 0 {
				continue
			}

			// Берется только 1-е значение (как и в SQL реализации)
			if !slices.ContainsFunc(values, func(value string) bool {
				return matchLike(attrFilter.Values[0], value)
			}) {
				return false
			}

		case core.BookFilterAttributeTypeIn:
			if len(attrFilter.Values) == 0 {
				continue
			}

			if !slices.ContainsFunc(values, func(value string) bool {
				return slices.Contains(attrFilter.Values, value)
			}) {
				return false
			}

		case core.BookFilterAttributeTypeCountEq:
			if len(values) != attrFilter.Count {
				return false
			}

		case core.BookFilterAttributeTypeCountGt:
			if len(values) <= attrFilter.Count {
				return false
			}

		case core.BookFilterAttributeTypeCountLt:
			// Случай "нет значений" ожидается только для 1, аналог NOT EXISTS в SQL
			if attrFilter.Count <= 1 {
				if len(values) > 0 {
					return false
				}
			} else if len(values) == 0 || len(values) >= attrFilter.Count {
				return false
			}

		case core.BookFilterAttributeTypeNone:
			continue

		default:
			continue
		}
	}

	return true
}

// parsedAttributeValues соберет уникальные значения атрибута разобранной книги,
// берутся сырые значения с внешнего источника.
func parsedAttributeValues(book hproxymodel.Book, code string) []string {
	for _, attr := range book.Attributes {
		if attr.Code == code {
			return pkg.Unique(attr.StringValues())
		}
	}

	return nil
}

// matchLike проверит совпадение строки с шаблоном SQL ILIKE:
// `_` - любой один символ, `%` - любая последовательность символов, регистр не учитывается.
// Шаблон дополняется маской `%с обеих сторон` (поиск подстроки), как и в SQL реализации.
func matchLike(pattern, value string) bool {
	var exp strings.Builder

	exp.Grow(len(pattern) + 16)
	exp.WriteString(`(?is).*`)

	for _, char := range pattern {
		switch char {
		case '_':
			exp.WriteString(".")
		case '%':
			exp.WriteString(".*")
		default:
			exp.WriteString(regexp.QuoteMeta(string(char)))
		}
	}

	exp.WriteString(".*")

	re, err := regexp.Compile(exp.String())
	if err == nil {
		return re.MatchString(value)
	}

	// Фолбэк на совпадение по подстроке
	return strings.Contains(strings.ToLower(value), strings.ToLower(pattern))
}
