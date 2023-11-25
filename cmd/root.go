/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/lucaschain/beholder/domain"
)

var rootCmd = &cobra.Command{
	Use:   "beholder [path] -- command-to-run",
	Short: "Simple command line file watcher",

	Long: `Beholder is a file watcher that can be used to run commands when files change.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var path = args[0]
		var command = strings.Join(args[1:], " ")

		fmt.Printf("Watching path: %s and running command: '%s'\n", path, command)

		domain.StartWatching(path, command)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var path string

func init() {
	// rootCmd.PersistentFlags().StringVar(&watchedPath, "path", "p", "config file (default is $HOME/.main.yaml)")
}
