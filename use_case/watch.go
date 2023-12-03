package use_case

import (
	"fmt"
	"log"

	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/core/event_types"
	"github.com/lucaschain/beholder/infrastructure"
)

func onFileChange(command []string, allowedTypes []string, allowFailing bool) core.ChangeCallback {
	return func(event *core.ChangeEvent, err *error) {
		if err != nil {
			log.Fatal(err)
		}

		if event_types.Filter(event.Type, allowedTypes) {
			command := core.CommandTokens(command, event)
			commandError := infrastructure.Command(command)

			if commandError != nil && !allowFailing {
				log.Fatal(commandError)
			}
		}
	}
}

func Watch(
	paths []string,
	command []string,
	allowedTypes []string,
	allowFailing bool,
) {
	fmt.Printf("Watching path: %s and running command: '%s'\n", paths, command)
	callback := onFileChange(command, allowedTypes, allowFailing)
	infrastructure.FileWatcher(paths, callback)
}
