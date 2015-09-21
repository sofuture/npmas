package version

import (
	"fmt"
)

const (
	Version           = "0.1.0"
	VersionPrerelease = "dev"
)

var (
	GitCommit string
)

func GetVersion() string {
	return fmt.Sprintf("npmas v%s%s-%s", Version, VersionPrerelease, GitCommit)
}
