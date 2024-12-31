package entities

import (
	"net/url"

	"github.com/google/uuid"
)

type DeduplicateArchiveResult struct {
	TargetBookID  uuid.UUID
	OriginBookURL *url.URL
	// Процент (0-1) вхождения архива в книгу
	EntryPercentage float64
	// Процент (0-1) вхождения книги в архив
	ReverseEntryPercentage float64
}

// FIXME: надо проверить и покрыть тестами,
// сейчас есть странная аналомалия что обе функции расчета выдают не 100%,
// при этом функция поиска страниц не находит разницу.
func EntryPercentageForPagesNew(current, target []PageWithHash) float64 {
	targetHashes := make(map[FileHash]struct{}, len(target))

	for _, p := range target {
		// targetHashes[p.Hash()]++ // Альтернативный вариант
		targetHashes[p.Hash()] = struct{}{}
	}

	var hits int

	count := len(targetHashes)

	for _, p := range current {
		_, ok := targetHashes[p.Hash()]
		if ok {
			// hits += targetHashes[p.Hash()] // Альтернативный вариант
			hits++
		}

		delete(targetHashes, p.Hash())
	}

	return float64(hits) / float64(count)
}

func EntryPercentageForPages(current, target []PageWithHash) float64 {
	targetHashes := make(map[FileHash]struct{}, len(target))

	for _, p := range target {
		targetHashes[p.Hash()] = struct{}{}
	}

	var hits int

	for _, p := range current {
		_, ok := targetHashes[p.Hash()]
		if ok {
			hits++
		}
	}

	return float64(hits) / float64(len(target))
}

type DeduplicateBookResult struct {
	TargetBook  Book
	PreviewPage Page
	// Процент (0-1) вхождения книги в целевую книгу
	EntryPercentage float64
	// Процент (0-1) вхождения целевой книги в книгу
	ReverseEntryPercentage float64
}

type BookPagesCompareResult struct {
	OriginBook        Book
	TargetBook        Book
	OriginPreviewPage Page
	TargetPreviewPage Page

	OriginPages []Page
	BothPages   []Page
	TargetPages []Page
}
