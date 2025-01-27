package entities

import (
	"math"

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
	Pages      []PageWithDeadHash
	Attributes []AttributeToWeb

	PreviewPage Page
	ParsedPages bool
	Tags        []string

	Size BookSize
}

func (book BookToWeb) PageDownloadPercent() float64 {
	downloadedPageCount := pkg.SliceReduce(book.Pages, func(v int, p PageWithDeadHash) int {
		if p.Downloaded {
			v++
		}

		return v
	})

	return math.Round(float64(downloadedPageCount)*10000/float64(len(book.Pages))) / 100
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
