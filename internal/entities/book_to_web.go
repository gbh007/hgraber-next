package entities

import (
	"math"

	"github.com/google/uuid"

	"hgnext/internal/pkg"
)

const PageNumberForPreview int = 1

type AttributeToWeb struct {
	Code   string
	Name   string
	Values []string
}

// TODO: подумать что делать с такими моделями
type BookToWeb struct {
	Book       Book
	Pages      []PreviewPage
	Attributes []AttributeToWeb

	PreviewPage PreviewPage
	ParsedPages bool
	Tags        []string

	Size BookSize
}

func (book BookToWeb) PageDownloadPercent() float64 {
	downloadedPageCount := pkg.SliceReduce(book.Pages, func(v int, p PreviewPage) int {
		if p.Downloaded {
			v++
		}

		return v
	})

	return math.Round(float64(downloadedPageCount)*10000/float64(len(book.Pages))) / 100
}

type PreviewPage struct {
	PageNumber  int
	Ext         string
	Downloaded  bool
	FileID      uuid.UUID
	FSID        *uuid.UUID
	HasDeadHash *bool
}

type BookListToWeb struct {
	Books []BookToWeb
	Pages []int

	Count int
}

type BookCompareResultToWeb struct {
	BookPagesCompareResult

	OriginAttributes []AttributeToWeb
	BothAttributes   []AttributeToWeb
	TargetAttributes []AttributeToWeb
}
