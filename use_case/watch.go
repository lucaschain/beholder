package use_case

import (
	"context"
	"fmt"

	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/core/event_types"
)

type WatchConfig struct {
	Paths        []string
	Command      []string
	AllowedTypes []event_types.EventType
	AllowFailing bool
}

type FileWatcher func(paths []string, callback core.ChangeCallback, ctx context.Context)
type CommandRunner func(command []string) error

func onFileChange(c WatchConfig, commandRunner CommandRunner) core.ChangeCallback {
	return func(event *core.ChangeEvent, err *error) *error {
		if err != nil {
			return err
		}

		if event_types.Filter(event.Type, c.AllowedTypes) {
			command := core.CommandTokens(c.Command, event)
			commandError := commandRunner(command)

			if commandError != nil && !c.AllowFailing {
				return &commandError
			}
		}
		return nil
	}
}

func Watch(
	c WatchConfig,
	fileWatcher FileWatcher,
	commandRunner CommandRunner,
) {
	fmt.Printf("Watching path: %s and running command: '%s'\n", c.Paths, c.Command)
	callback := onFileChange(c, commandRunner)
	fileWatcher(c.Paths, callback, context.Background())
}
