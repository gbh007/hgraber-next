//go:build linux

package localfiles

import "golang.org/x/sys/unix"

func getAvailableSize(p string) int64 {
	var stat unix.Statfs_t

	_ = unix.Statfs(p, &stat)

	return int64(stat.Bavail * uint64(stat.Bsize)) //nolint:gosec // не возможный кейс
}
