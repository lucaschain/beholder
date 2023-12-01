package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var Version string = "dev"

var rootCmd = &cobra.Command{
	Use:     "beholder [path] -- command-to-run",
	Short:   "Simple command line file watcher",
	Version: Version,
	Long:    `Beholder is a file watcher that can be used to run commands when files change.`,
	Args:    cobra.MinimumNArgs(2),
	Run:     Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
