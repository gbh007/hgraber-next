package entities

import (
	"math"
	"time"

	"hgnext/internal/pkg"
)

type BookFilterShowType byte

const (
	BookFilterShowTypeAll = iota
	BookFilterShowTypeOnly
	BookFilterShowTypeExcept
)

type BookFilterOrderBy byte

const (
	BookFilterOrderByCreated = iota
	BookFilterOrderByName
	BookFilterOrderByID
	BookFilterOrderByPageCount
)

type BookFilterFields struct {
	Name string
}

type BookFilter struct {
	Limit  int
	Offset int

	OrderBy BookFilterOrderBy
	Desc    bool

	From time.Time
	To   time.Time

	OriginAttributes bool

	ShowDeleted    BookFilterShowType
	ShowVerified   BookFilterShowType
	ShowDownloaded BookFilterShowType

	Fields BookFilterFields
}

func (f *BookFilter) FillLimits(page, count int) {
	if page < 1 {
		page = 1
	}

	f.Offset = (page - 1) * count
	f.Limit = count
}

type AttributeToWeb struct {
	Code   string
	Name   string
	Values []string
}

type BookToWeb struct {
	Book       Book
	Pages      []Page
	Attributes []AttributeToWeb

	PreviewPage Page
	ParsedPages bool
	Tags        []string
	HasMoreTags bool
}

func (book BookToWeb) PageDownloadPercent() float64 {
	downloadedPageCount := pkg.SliceReduce(book.Pages, func(v int, p Page) int {
		if p.Downloaded {
			v++
		}

		return v
	})

	return math.Round(float64(downloadedPageCount)*10000/float64(len(book.Pages))) / 100
}

func (book BookToWeb) ParsedName() bool {
	return book.Book.Name != ""
}

type BookListToWeb struct {
	Books []BookToWeb
	Pages []int

	Count int
}
