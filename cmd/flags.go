package cmd

import (
	"fmt"

	"github.com/lucaschain/beholder/core/event_types"
	"github.com/spf13/cobra"
)

var types []string
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

	cmd.Flags().BoolVarP(
		&allowFailing,
		"allow-failing",
		"f",
		false,
		"Keep running when command fails",
	)
}
