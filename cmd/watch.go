package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/infrastructure"
	"github.com/spf13/cobra"
)

func onFileChange(command []string) core.Callback {
	return func(event *fsnotify.Event, err *error) {
		if err != nil {
			log.Fatal(err)
		}
		if event.Has(fsnotify.Write) {
			commandError := infrastructure.RunCommand(core.Replace(command, event))

			if commandError != nil {
				log.Fatal(commandError)
			}
		}
	}
}

func Run(cmd *cobra.Command, args []string) {
	var command = strings.Join(args, " ")
	fmt.Printf("Watching path: %s and running command: '%s'\n", path, command)
	core.Start(path, onFileChange(args))
}
