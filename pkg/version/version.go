// Package version provides the version of this tool.
// Ref: https://pkg.go.dev/embed
package version

import (
	_ "embed"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string
)
