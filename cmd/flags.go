package cmd

import (
	"fmt"

	"github.com/lucaschain/beholder/core/event_types"
	"github.com/spf13/cobra"
)

var types []string
var extensions []string
var allowFailing bool
var defaultTypes = []string{"WRITE"}

func SetFlags(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(
		&types,
		"type",
		"t",
		defaultTypes,
		fmt.Sprintf("Event types to watch, options: %s", event_types.EventTypes),
	)

	cmd.Flags().StringSliceVarP(
		&extensions,
		"extension",
		"e",
		[]string{},
		`If provided, only files with the given extensions will trigger the command. E.g.:
$ beholder . -e go -e yml -- go test ./...`,
	)

	cmd.Flags().BoolVarP(
		&allowFailing,
		"allow-failing",
		"f",
		true,
		"Keep running when command fails",
	)
}
