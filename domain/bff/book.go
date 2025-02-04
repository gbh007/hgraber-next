package bff

import (
	"math"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

type StatusFlag byte

const (
	UnknownStatusFlag StatusFlag = iota
	TrueStatusFlag
	FalseStatusFlag
)

func NewStatusFlag(ok bool) StatusFlag {
	if ok {
		return TrueStatusFlag
	}

	return FalseStatusFlag
}

type BookDetails struct {
	Book       core.Book
	Pages      []PreviewPage
	Attributes []AttributeToWeb

	PreviewPage PreviewPage

	Size core.BookSize

	FSDisposition []BookDetailsFSDisposition
}

func (book BookDetails) PageDownloadPercent() float64 {
	downloadedPageCount := pkg.SliceReduce(book.Pages, func(v int, p PreviewPage) int {
		if p.Downloaded {
			v++
		}

		return v
	})

	return math.Round(float64(downloadedPageCount)*10000/float64(len(book.Pages))) / 100
}

func (book BookDetails) AvgPageSize() int64 {
	if len(book.Pages) == 0 {
		return 0
	}

	return book.Size.Total / int64(len(book.Pages))
}

type BookDetailsFSDisposition struct {
	ID   uuid.UUID
	Name string
	core.SizeWithCount
}

type BookShort struct {
	Book        core.Book
	PreviewPage PreviewPage
	Tags        []string
}

type BookList struct {
	Books []BookShort
	Pages []int

	Count int
}

type AttributeToWeb struct {
	Code   string
	Name   string
	Values []string
}

// TODO: подумать что делать с такими моделями
type BookWithPreviewPage struct {
	core.Book
	PreviewPage PreviewPage
}
