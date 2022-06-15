package version

import "fmt"

var Version string
var Commit string

func GetVersion() string {
	if Version == "" {
		Version = "0.0.0-dev"
	}
	return Version
}

func GetCommit() string {
	if Commit == "" {
		Commit = "unknown_commit"
	}
	return Commit
}

func GetVersionWithCommit() string {
	return fmt.Sprintf("%s-%s", GetVersion(), GetCommit())
}
