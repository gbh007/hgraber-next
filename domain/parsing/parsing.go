package parsing

import (
	"net/url"

	"github.com/google/uuid"
)

type ParseFlags struct {
	AutoVerify bool
	ReadOnly   bool
	// TODO: временное размещение
	MaxAuthors int
}

type FirstHandleMultipleResult struct {
	TotalCount     int
	LoadedCount    int
	DuplicateCount int
	SkippedCount   int
	ErrorCount     int
	NotHandled     []url.URL
	Details        []BookHandleResult
}

func (result *FirstHandleMultipleResult) RegisterError(u url.URL, reason string) {
	result.ErrorCount++
	result.TotalCount++
	result.NotHandled = append(result.NotHandled, u)
	result.Details = append(result.Details, BookHandleResult{
		URL:         u,
		ErrorReason: reason,
	})
}

func (result *FirstHandleMultipleResult) RegisterDuplicate(u url.URL, ids []uuid.UUID) {
	result.TotalCount++
	result.DuplicateCount++
	result.Details = append(result.Details, BookHandleResult{
		URL:          u,
		IsDuplicate:  true,
		DuplicateIDs: ids,
	})
}

func (result *FirstHandleMultipleResult) RegisterSkipped(u url.URL) {
	result.TotalCount++
	result.SkippedCount++
	result.Details = append(result.Details, BookHandleResult{
		URL:       u,
		IsSkipped: true,
	})
}

func (result *FirstHandleMultipleResult) RegisterHandled(u url.URL, id uuid.UUID) {
	result.TotalCount++
	result.LoadedCount++
	result.Details = append(result.Details, BookHandleResult{
		URL:       u,
		IsHandled: true,
		ID:        id,
	})
}
