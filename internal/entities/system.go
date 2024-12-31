package entities

type SystemSizeInfo struct {
	BookCount            int
	DownloadedBookCount  int
	VerifiedBookCount    int
	BookUnparsedCount    int
	PageCount            int
	PageUnloadedCount    int
	PageWithoutBodyCount int
	PageFileSize         int64
	FileSize             int64
}

type SystemSizeInfoWithMonitor struct {
	SystemSizeInfo
	Workers []SystemWorkerStat
}

type SystemWorkerStat struct {
	Name         string
	InQueueCount int
	InWorkCount  int
	RunnersCount int
}
