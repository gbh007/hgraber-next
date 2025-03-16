package bff

import "github.com/gbh007/hgraber-next/domain/core"

type DeduplicateBookResult struct {
	TargetBook  core.Book
	PreviewPage PreviewPage
	// Процент (0-1) вхождения книги в целевую книгу
	EntryPercentage float64
	// Процент (0-1) вхождения целевой книги в книгу
	ReverseEntryPercentage float64
	// Процент (0-1) вхождения архива в книгу без учета мертвых хешей
	EntryPercentageWithoutDeadHashes float64
	// Процент (0-1) вхождения книги в архив без учета мертвых хешей
	ReverseEntryPercentageWithoutDeadHashes float64

	SharedSize                  int64
	SharedSizeWithoutDeadHashes int64

	SharedPages                  int
	SharedPagesWithoutDeadHashes int

	TargetSize core.SizeWithCount
}

type BookPagesCompareResult struct {
	OriginBook core.Book
	TargetBook core.Book

	OriginSize core.SizeWithCount
	TargetSize core.SizeWithCount

	OriginPreviewPage PreviewPage
	TargetPreviewPage PreviewPage

	OriginPages []PreviewPage
	BothPages   []PreviewPage
	TargetPages []PreviewPage

	// Процент (0-1) вхождения книги в целевую книгу
	EntryPercentage float64
	// Процент (0-1) вхождения целевой книги в книгу
	ReverseEntryPercentage float64
	// Процент (0-1) вхождения архива в книгу без учета мертвых хешей
	EntryPercentageWithoutDeadHashes float64
	// Процент (0-1) вхождения книги в архив без учета мертвых хешей
	ReverseEntryPercentageWithoutDeadHashes float64
}

type BookAttributesCompareResult struct {
	OriginAttributes map[string][]string
	BothAttributes   map[string][]string
	TargetAttributes map[string][]string
}

type BookCompareResult struct {
	BookPagesCompareResult

	OriginAttributes []AttributeToWeb
	BothAttributes   []AttributeToWeb
	TargetAttributes []AttributeToWeb
}
