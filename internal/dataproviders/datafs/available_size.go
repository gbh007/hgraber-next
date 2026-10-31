//go:build !linux

package datafs

func getAvailableSize(_ string) int64 { return 0 }
