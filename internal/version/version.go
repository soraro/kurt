package version

import (
	"fmt"
	"runtime"
)

// NOTE: these variables are injected at build time
var (
	version           string = "development"
	gitSHA, buildTime string
	build             Build
)

type Build struct {
	Version   string `json:"version,omitempty"`
	GitSHA    string `json:"git,omitempty"`
	BuildTime string `json:"buildTime,omitempty"`
	GoVersion string `json:"goversion,omitempty"`
}

func initBuild() {
	build.Version = version
	if len(gitSHA) >= 7 {
		build.GitSHA = gitSHA[:7]
	}
	build.BuildTime = buildTime
	build.GoVersion = runtime.Version()
}

func Version() string {
	initBuild()
	return fmt.Sprintf("%s\n%s", build.Version, build)
}
