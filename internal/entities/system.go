package entities

type SystemSizeInfo struct {
	BookCount         int
	BookUnparsedCount int
	PageCount         int
	PageUnloadedCount int
	PageFileSize      int64
}
