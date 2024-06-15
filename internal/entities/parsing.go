package entities

import "net/url"

type FirstHandleMultipleResult struct {
	TotalCount     int64
	LoadedCount    int64
	DuplicateCount int64
	ErrorCount     int64
	NotHandled     []url.URL
	Details        []BookHandleResult
}

type BookHandleResult struct {
	URL         url.URL
	IsDuplicate bool
	IsHandled   bool
	ErrorReason string
}
