package entities

import (
	"math"

	"github.com/google/uuid"

	"hgnext/internal/pkg"
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

type BFFBookDetails struct {
	Book       Book
	Pages      []BFFPreviewPage
	Attributes []AttributeToWeb

	PreviewPage BFFPreviewPage

	Size BookSize
}

func (book BFFBookDetails) PageDownloadPercent() float64 {
	downloadedPageCount := pkg.SliceReduce(book.Pages, func(v int, p BFFPreviewPage) int {
		if p.Downloaded {
			v++
		}

		return v
	})

	return math.Round(float64(downloadedPageCount)*10000/float64(len(book.Pages))) / 100
}

type BFFBookShort struct {
	Book        Book
	PreviewPage BFFPreviewPage
	Tags        []string
}

type BFFBookList struct {
	Books []BFFBookShort
	Pages []int

	Count int
}

const PageNumberForPreview int = 1

type AttributeToWeb struct {
	Code   string
	Name   string
	Values []string
}

type BFFPreviewPage struct {
	PageNumber  int
	Ext         string
	Downloaded  bool
	FileID      uuid.UUID
	FSID        uuid.UUID
	HasDeadHash StatusFlag
}

type BookCompareResultToWeb struct {
	BookPagesCompareResult

	OriginAttributes []AttributeToWeb
	BothAttributes   []AttributeToWeb
	TargetAttributes []AttributeToWeb
}
