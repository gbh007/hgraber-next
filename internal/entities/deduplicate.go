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
