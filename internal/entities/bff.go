package entities

import (
	"math"

	"hgnext/internal/pkg"
)

type BFFBookDetails struct {
	Book       Book
	Pages      []PreviewPage
	Attributes []AttributeToWeb

	PreviewPage PreviewPage

	Size BookSize
}

func (book BFFBookDetails) PageDownloadPercent() float64 {
	downloadedPageCount := pkg.SliceReduce(book.Pages, func(v int, p PreviewPage) int {
		if p.Downloaded {
			v++
		}

		return v
	})

	return math.Round(float64(downloadedPageCount)*10000/float64(len(book.Pages))) / 100
}
