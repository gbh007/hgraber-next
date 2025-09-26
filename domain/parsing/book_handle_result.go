package parsing

import (
	"net/url"

	"github.com/google/uuid"
)

type BookHandleResult struct {
	URL url.URL

	IsDuplicate  bool
	DuplicateIDs []uuid.UUID

	IsHandled bool
	ID        uuid.UUID

	ErrorReason string
}

type MultiHandleMultipleResult struct {
	TotalCount  int
	LoadedCount int
	ErrorCount  int
	NotHandled  []url.URL
	Details     FirstHandleMultipleResult
}

func (result *MultiHandleMultipleResult) RegisterError(u url.URL, reason string) {
	result.ErrorCount++
	result.TotalCount++
	result.NotHandled = append(result.NotHandled, u)
}

func (result *MultiHandleMultipleResult) RegisterHandled(u url.URL) {
	result.TotalCount++
	result.LoadedCount++
}
