package version

// Info describes the current version of the code being run.
type Info struct {
	Revision string
	Version  string
}

// Revision is the git revision that was compiled
var Revision = "unknown revision"

// Version number that is being run at the moment. Version should use semver.
var Version = "unknown version"
