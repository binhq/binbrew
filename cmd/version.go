package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Provisioned by ldflags
var (
	Version    string
	CommitHash string
	BuildDate  string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("binbrew version %s, build %s (at %s)\n", Version, CommitHash, BuildDate)
	},
}
