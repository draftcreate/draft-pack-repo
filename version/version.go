package version

import (
	"fmt"
)

// Version contains the semver release, git commit, and git tree state.
type Version struct {
	SemVer       string `json:"semver"`
	GitCommit    string `json:"git-commit"`
	GitTreeState string `json:"git-tree-state"`
}

func (v *Version) String() string {
	ver := v.SemVer
	// show commit metadata if this is an unofficial release
	if Release == "canary" && GitCommit != "" {
		ver = fmt.Sprintf("%s+%s.%s", v.SemVer, v.GitCommit, v.GitTreeState)
	}
	return ver
}

var (
	// Release is the current release of pack-man.
	// The release is of the format Major.Minor.Patch[-Prerelease][+BuildMetadata]
	//
	// If it is a development build, the release name is called "canary".
	//
	// This number is incremented automatically via build.sh.
	Release = "v0.2.0"

	// BuildMetadata is extra build time data
	BuildMetadata = ""
	// GitCommit is the git sha1
	GitCommit = ""
	// GitTreeState is the state of the git tree
	GitTreeState = ""
)

// getVersion returns the semver string of the version
func getVersion() string {
	if BuildMetadata == "" {
		return Release
	}
	return Release + "+" + BuildMetadata
}

// New returns the semver interpretation of the version.
func New() *Version {
	return &Version{
		SemVer:       getVersion(),
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
	}
}
