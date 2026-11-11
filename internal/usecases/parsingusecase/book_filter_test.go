package parsingusecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gbh007/hgraber-next/internal/domain/core"
	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
)

func TestBookMatchesFilter(t *testing.T) {
	testCases := []struct {
		desc   string
		filter core.BookFilter
		book   hproxymodel.Book
		want   bool
	}{
		{
			desc: "пустой фильтр пропускает все",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Name:       "",
					Attributes: []core.BookFilterAttribute{},
				},
			},
			book: hproxymodel.Book{
				Name: "любое имя",
			},
			want: true,
		},
		{
			desc: "фильтр по имени подстрока",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Name: "мир",
				},
			},
			book: hproxymodel.Book{
				Name: "В МИРЕ ЖИВОТНЫХ",
			},
			want: true,
		},
		{
			desc: "фильтр по имени не совпал",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Name: "мир",
				},
			},
			book: hproxymodel.Book{
				Name: "что-то другое",
			},
			want: false,
		},
		{
			desc: "фильтр по имени спецсимволы SQL LIKE (%)",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Name: "ск%ж",
				},
			},
			book: hproxymodel.Book{
				Name: "СКАЖЖ",
			},
			want: true,
		},
		{
			desc: "фильтр по имени спецсимволы SQL LIKE (_)",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Name: "ск_ж",
				},
			},
			book: hproxymodel.Book{
				Name: "скаж",
			},
			want: true,
		},
		{
			desc: "фильтр по имени спецсимволы SQL LIKE 3",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Name: "ск_ж",
				},
			},
			book: hproxymodel.Book{
				Name: "скж",
			},
			want: false,
		},
		{
			desc: "атрибут IN любое из",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeAuthor,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"автор 1", "автор 2"},
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "автор 2"},
						},
					},
				},
			},
			want: true,
		},
		{
			desc: "атрибут IN точное совпадение",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeAuthor,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"автор 1"},
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "автор 11"},
						},
					},
				},
			},
			want: false,
		},
		{
			desc: "атрибут отсутствует",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeAuthor,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"автор 1"},
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeTag,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "тег"},
						},
					},
				},
			},
			want: false,
		},
		{
			desc: "атрибут LIKE подстрока",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeTag,
							Type:   core.BookFilterAttributeTypeLike,
							Values: []string{"фанта"},
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeTag,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "научная фантастика"},
						},
					},
				},
			},
			want: true,
		},
		{
			desc: "несколько атрибутов через И",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeAuthor,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"автор 1"},
						},
						{
							Code:   core.AttributeCodeTag,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"тег"},
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "автор 1"},
						},
					},
					{
						Code: core.AttributeCodeTag,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "тег"},
						},
					},
				},
			},
			want: true,
		},
		{
			desc: "несколько атрибутов через И, второй не совпал",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:   core.AttributeCodeAuthor,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"автор 1"},
						},
						{
							Code:   core.AttributeCodeTag,
							Type:   core.BookFilterAttributeTypeIn,
							Values: []string{"тег"},
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "автор 1"},
						},
					},
				},
			},
			want: false,
		},
		{
			desc: "count_eq",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:  core.AttributeCodeAuthor,
							Type:  core.BookFilterAttributeTypeCountEq,
							Count: 2,
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "а1"},
							{ExtName: "а2"},
						},
					},
				},
			},
			want: true,
		},
		{
			desc: "count_eq считает уникальные значения",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:  core.AttributeCodeAuthor,
							Type:  core.BookFilterAttributeTypeCountEq,
							Count: 2,
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "а1"},
							{ExtName: "а1"},
							{Name: "а2"},
						},
					},
				},
			},
			want: true,
		},
		{
			desc: "count_gt",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:  core.AttributeCodeAuthor,
							Type:  core.BookFilterAttributeTypeCountGt,
							Count: 1,
						},
					},
				},
			},
			book: hproxymodel.Book{
				Attributes: []hproxymodel.BookAttribute{
					{
						Code: core.AttributeCodeAuthor,
						Values: []hproxymodel.BookAttributeValue{
							{ExtName: "а1"},
						},
					},
				},
			},
			want: false,
		},
		{
			desc: "count_lt требует наличия значения",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:  core.AttributeCodeAuthor,
							Type:  core.BookFilterAttributeTypeCountLt,
							Count: 3,
						},
					},
				},
			},
			book: hproxymodel.Book{},
			want: false,
		},
		{
			desc: "count_lt 1 ожидает отсутствие значений",
			filter: core.BookFilter{
				Fields: core.BookFilterFields{
					Attributes: []core.BookFilterAttribute{
						{
							Code:  core.AttributeCodeAuthor,
							Type:  core.BookFilterAttributeTypeCountLt,
							Count: 1,
						},
					},
				},
			},
			book: hproxymodel.Book{},
			want: true,
		},
		{
			desc: "page_count от совпало",
			filter: core.BookFilter{
				PageCountFrom: 10,
			},
			book: hproxymodel.Book{
				PageCount: 10,
			},
			want: true,
		},
		{
			desc: "page_count от меньше",
			filter: core.BookFilter{
				PageCountFrom: 10,
			},
			book: hproxymodel.Book{
				PageCount: 9,
			},
			want: false,
		},
		{
			desc: "page_count до совпало",
			filter: core.BookFilter{
				PageCountTo: 10,
			},
			book: hproxymodel.Book{
				PageCount: 9,
			},
			want: true,
		},
		{
			desc: "page_count до не совпало",
			filter: core.BookFilter{
				PageCountTo: 10,
			},
			book: hproxymodel.Book{
				PageCount: 10,
			},
			want: false,
		},
		{
			desc: "page_count диапазон внутри",
			filter: core.BookFilter{
				PageCountFrom: 10,
				PageCountTo:   100,
			},
			book: hproxymodel.Book{
				PageCount: 50,
			},
			want: true,
		},
		{
			desc: "page_count диапазон, нет данных",
			filter: core.BookFilter{
				PageCountFrom: 10,
			},
			book: hproxymodel.Book{},
			want: false,
		},
		{
			desc: "page_count диапазон и имя, имя не совпало",
			filter: core.BookFilter{
				PageCountFrom: 10,
				Fields: core.BookFilterFields{
					Name: "мир",
				},
			},
			book: hproxymodel.Book{
				Name:      "что-то другое",
				PageCount: 50,
			},
			want: false,
		},
		{
			desc: "page_count диапазон и имя, оба совпали",
			filter: core.BookFilter{
				PageCountFrom: 10,
				Fields: core.BookFilterFields{
					Name: "мир",
				},
			},
			book: hproxymodel.Book{
				Name:      "в мире животных",
				PageCount: 50,
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.want, bookMatchesFilter(tc.filter, tc.book), "bookMatchesFilter")
		})
	}
}
