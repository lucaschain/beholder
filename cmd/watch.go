package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/core/event_types"
	"github.com/lucaschain/beholder/infrastructure"
	"github.com/spf13/cobra"
)

func onFileChange(command []string) core.ChangeCallback {
	return func(event *core.ChangeEvent, err *error) {
		if err != nil {
			log.Fatal(err)
		}

		if event.Type == event_types.Write {
			command := core.CommandTokens(command, event)
			commandError := infrastructure.Command(command)

			if commandError != nil {
				log.Fatal(commandError)
			}
		}
	}
}

func Run(cmd *cobra.Command, args []string) {
	paths := strings.Split(args[0], ",")
	var command = strings.Join(args[1:], " ")
	fmt.Printf("Watching path: %s and running command: '%s'\n", paths, command)
	infrastructure.FileWatcher(paths, onFileChange(args[1:]))
}
