package entities

import "net/url"

type FirstHandleMultipleResult struct {
	TotalCount     int
	LoadedCount    int
	DuplicateCount int
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

func (result *FirstHandleMultipleResult) RegisterDuplicate(u url.URL) {
	result.TotalCount++
	result.DuplicateCount++
	result.Details = append(result.Details, BookHandleResult{
		URL:         u,
		IsDuplicate: true,
	})
}

func (result *FirstHandleMultipleResult) RegisterHandled(u url.URL) {
	result.TotalCount++
	result.LoadedCount++
	result.Details = append(result.Details, BookHandleResult{
		URL:       u,
		IsHandled: true,
	})
}

type BookHandleResult struct {
	URL         url.URL
	IsDuplicate bool
	IsHandled   bool
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
