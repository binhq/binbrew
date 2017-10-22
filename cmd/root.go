package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var logger = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "binbrew",
	Short: "The binary package manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logger.SetLevel(logrus.DebugLevel)
		} else if verbose, err := strconv.ParseBool(os.Getenv("BINBREW_VERBOSE")); err == nil && verbose {
			logger.SetLevel(logrus.DebugLevel)
		}
	},
}

func init() {
	if filepath.Base(os.Args[0]) == "bin" {
		rootCmd.Use = "bin"
	}
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")

	rootCmd.AddCommand(versionCmd, getCmd)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
