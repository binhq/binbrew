package cmd

import (
	"os"

	"github.com/binhq/binbrew/pkg"
	"github.com/binhq/binbrew/pkg/github"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [binary] [path]",
	Short: "Download a binary to a given path (or current directory)",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		resolver := &pkg.Resolver{
			Providers: map[string]pkg.Provider{
				"github": &github.Provider{},
			},
		}

		logger.WithField("binary", args[0]).Debug("resolving binary")

		binary, err := resolver.Resolve(args[0])
		if err != nil {
			logger.Error(err)
			return
		}

		downloader := pkg.NewDownloader(pkg.NewCache(), logger)

		var dst string
		if len(args) == 2 {
			dst = args[1]
		} else if wd, err := os.Getwd(); err == nil {
			dst = wd
		} else {
			dst = "."
		}

		err = downloader.Download(binary, dst)
		if err != nil {
			logger.Error(err)
			return
		}
	},
}
