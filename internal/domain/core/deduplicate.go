package core

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

// TODO: надо проверить и покрыть тестами.
func EntryPercentageForPages(current, target []PageWithHash, deadHashes map[FileHash]struct{}) float64 {
	targetHashes := make(map[FileHash]struct{}, len(target))

	for _, p := range target {
		if _, ok := deadHashes[p.FileHash]; ok {
			continue
		}

		targetHashes[p.FileHash] = struct{}{}
	}

	var hits int

	count := len(targetHashes)

	for _, p := range current {
		if _, ok := deadHashes[p.FileHash]; ok {
			continue
		}

		if _, ok := targetHashes[p.FileHash]; ok {
			hits++

			delete(targetHashes, p.FileHash)
		}
	}

	switch { // Обработка крайних случаев
	case hits == 0 && count == 0:
		return 1

	case hits == 0:
		return 0

	case count == 0:
		return 1
	}

	return float64(hits) / float64(count)
}
