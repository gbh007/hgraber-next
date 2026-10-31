//go:build !linux

package localfs

func getAvailableSize(_ string) int64 { return 0 }
