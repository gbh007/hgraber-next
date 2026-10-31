//go:build !linux

package serverfs

func getAvailableSize(_ string) int64 { return 0 }
