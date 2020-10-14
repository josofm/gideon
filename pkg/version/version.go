package version

import "fmt"

// This string will be overwritten during the build process.
//nolint:gochecknoglobals
var (
	// Semver semantic version
	Semver = ""
	// GitSHA git commit sha
	GitSHA = ""
)

// Version returns a newline-terminated string describing the current
// version of the build.
func Version() string {
	if GitSHA == "" {
		return fmt.Sprintf("Version: %s\nGit hash: devel\n", Semver)
	}

	return fmt.Sprintf("Version: %s\nGit hash: %s\n", Semver, GitSHA)
}
