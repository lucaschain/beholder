package cmd

import (
	"os"
	"strings"

	"github.com/lucaschain/beholder/core/event_types"
	"github.com/lucaschain/beholder/infrastructure"
	"github.com/lucaschain/beholder/use_case"
	"github.com/spf13/cobra"
)

var Version string = "dev"

var rootCmd = &cobra.Command{
	Use:     "beholder [path] -- command-to-run",
	Short:   "Simple command line file watcher",
	Version: Version,
	Long:    `Beholder is a language-agnostic file watcher that can be used to run commands when files change.`,
	Args:    cobra.MinimumNArgs(2),
	Run:     Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	SetFlags(rootCmd)
}

func Run(cmd *cobra.Command, args []string) {
	paths := strings.Split(args[0], ",")
	command := args[1:]

	var allowedTypes []event_types.EventType

	for _, t := range types {
		allowedTypes = append(allowedTypes, event_types.FromString(t))
	}

	watchConfig := use_case.WatchConfig{
		Paths:        paths,
		Command:      command,
		AllowedTypes: allowedTypes,
		Extensions:   extensions,
		AllowFailing: allowFailing,
	}
	use_case.Watch(watchConfig, infrastructure.FileWatcher, infrastructure.Command)
}
