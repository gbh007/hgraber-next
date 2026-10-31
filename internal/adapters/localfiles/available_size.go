//go:build !linux

package localfiles

func getAvailableSize(_ string) int64 { return 0 }
