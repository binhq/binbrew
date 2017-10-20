package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	quiet bool
)

var rootCmd = &cobra.Command{
	Use:   "binbrew",
	Short: "The binary package manager",
}

func init() {
	if os.Args[0] == "bin" {
		rootCmd.Use = "bin"
	}
	//rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Do not show log output")

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
