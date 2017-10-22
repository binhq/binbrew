package cmd

import (
	"fmt"
	"os"

	"github.com/binhq/binbrew/pkg"
	"github.com/binhq/binbrew/pkg/github"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [binary] [path]",
	Short: "Download a binary to a given path (or current directory)",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		resolver := &pkg.Resolver{
			Providers: map[string]pkg.Provider{
				"github": &github.Provider{},
			},
		}

		binary, err := resolver.Resolve(args[0])
		if err != nil {
			return err
		}

		downloader := pkg.NewDownloader(pkg.NewCache())

		var dst string
		if len(args) == 2 {
			dst = args[1]
		} else if wd, err := os.Getwd(); err == nil {
			dst = wd
		} else {
			dst = "."
		}

		return downloader.Download(binary, dst)
	},
}
