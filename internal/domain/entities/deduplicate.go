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
