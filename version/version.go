package version

import (
	"runtime"
	"runtime/debug"
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

func init() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		Version = info.Main.Version
	}

	for _, v := range info.Settings {
		switch v.Key {
		case "vcs.revision":
			Commit = v.Value
		case "vcs.time":
			BuildAt = v.Value
		}
	}
}
