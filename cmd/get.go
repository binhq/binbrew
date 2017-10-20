package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a binary to a given path (or current directory)",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}
