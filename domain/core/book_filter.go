//nolint:decorder // будет исправлено позднее
package core

import "time"

type BookFilterShowType byte

const (
	BookFilterShowTypeAll BookFilterShowType = iota
	BookFilterShowTypeOnly
	BookFilterShowTypeExcept
)

type BookFilterOrderBy byte

const (
	BookFilterOrderByCreated BookFilterOrderBy = iota
	BookFilterOrderByName
	BookFilterOrderByID
	BookFilterOrderByPageCount
	BookFilterOrderByCalcPageCount
	BookFilterOrderByCalcFileCount
	BookFilterOrderByCalcDeadHashCount
	BookFilterOrderByCalcPageSize
	BookFilterOrderByCalcFileSize
	BookFilterOrderByCalcDeadHashSize
	BookFilterOrderByCalculatedAt
)

type BookFilterAttributeType byte

const (
	BookFilterAttributeTypeNone BookFilterAttributeType = iota
	BookFilterAttributeTypeLike
	BookFilterAttributeTypeIn
	BookFilterAttributeTypeCountEq
	BookFilterAttributeTypeCountGt
	BookFilterAttributeTypeCountLt
)

type BookFilterLabelType byte

const (
	BookFilterLabelTypeNone BookFilterLabelType = iota
	BookFilterLabelTypeLike
	BookFilterLabelTypeIn
	BookFilterLabelTypeCountEq
	BookFilterLabelTypeCountGt
	BookFilterLabelTypeCountLt
)

type BookFilterFields struct {
	Name       string
	Attributes []BookFilterAttribute
	Labels     []BookFilterLabel
}

type BookFilterAttribute struct {
	Code   string
	Type   BookFilterAttributeType
	Values []string
	Count  int
}

type BookFilterLabel struct {
	Name   string
	Type   BookFilterLabelType
	Values []string
	Count  int
}

type BookFilter struct {
	Limit  int
	Offset int

	OrderBy BookFilterOrderBy
	Desc    bool

	From time.Time
	To   time.Time

	OriginAttributes bool

	ShowDeleted        BookFilterShowType
	ShowVerified       BookFilterShowType
	ShowDownloaded     BookFilterShowType
	ShowRebuilded      BookFilterShowType
	ShowWithoutPages   BookFilterShowType
	ShowWithoutPreview BookFilterShowType

	Fields BookFilterFields
}

func (f *BookFilter) FillLimits(page, count int) {
	if page < 1 {
		page = 1
	}

	f.Offset = (page - 1) * count
	f.Limit = count
}
