package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var path string

var rootCmd = &cobra.Command{
	Use:   "beholder [path] -- command-to-run",
	Short: "Simple command line file watcher",

	Long: `Beholder is a file watcher that can be used to run commands when files change.`,
	Args: cobra.MinimumNArgs(1),
	Run:  Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&path, "path", "p", "Path to watch")
}
