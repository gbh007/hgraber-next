//go:build !linux

package localFiles

func getAvailableSize(_ string) int64 { return 0 }
