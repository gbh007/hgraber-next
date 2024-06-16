package entities

type SystemSizeInfo struct {
	BookCount         int
	BookUnparsedCount int
	PageCount         int
	PageUnloadedCount int
	PageFileSize      int64
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
