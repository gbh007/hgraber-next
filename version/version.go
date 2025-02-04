package version

import (
	"runtime"
)

var (
	GoVersion = runtime.Version()
	GoOS      = runtime.GOOS
	GoArch    = runtime.GOARCH
)

var (
	Version string
	Commit  string
	Branch  string
	BuildAt string
)
