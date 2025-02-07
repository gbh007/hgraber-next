package systemmodel

import "github.com/google/uuid"

type SystemSizeInfo struct {
	BookCount           int
	DownloadedBookCount int
	VerifiedBookCount   int
	RebuildedBookCount  int
	BookUnparsedCount   int
	DeletedBookCount    int

	PageCount            int
	PageUnloadedCount    int
	PageWithoutBodyCount int
	DeletedPageCount     int

	DeadHashCount int

	FileCountByFS         map[uuid.UUID]int64
	UnhashedFileCountByFS map[uuid.UUID]int64
	InvalidFileCountByFS  map[uuid.UUID]int64
	DetachedFileCountByFS map[uuid.UUID]int64

	PageFileSizeByFS map[uuid.UUID]int64
	FileSizeByFS     map[uuid.UUID]int64
}

func (info SystemSizeInfo) FileCountByFSSum() int64 {
	var s int64

	for _, v := range info.FileCountByFS {
		s += v
	}

	return s
}

func (info SystemSizeInfo) UnhashedFileCountByFSSum() int64 {
	var s int64

	for _, v := range info.UnhashedFileCountByFS {
		s += v
	}

	return s
}

func (info SystemSizeInfo) InvalidFileCountByFSSum() int64 {
	var s int64

	for _, v := range info.InvalidFileCountByFS {
		s += v
	}

	return s
}

func (info SystemSizeInfo) DetachedFileCountByFSSum() int64 {
	var s int64

	for _, v := range info.DetachedFileCountByFS {
		s += v
	}

	return s
}

func (info SystemSizeInfo) PageFileSizeByFSSum() int64 {
	var s int64

	for _, v := range info.PageFileSizeByFS {
		s += v
	}

	return s
}

func (info SystemSizeInfo) FileSizeByFSSum() int64 {
	var s int64

	for _, v := range info.FileSizeByFS {
		s += v
	}

	return s
}

type SystemWorkerStat struct {
	Name         string
	InQueueCount int
	InWorkCount  int
	RunnersCount int
}
