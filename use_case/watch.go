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
	Extensions   []string
	AllowFailing bool
}

type FileWatcher func(paths []string, callback core.ChangeCallback, ctx context.Context) *error
type CommandRunner func(command []string) error

func onFileChange(c WatchConfig, commandRunner CommandRunner) core.ChangeCallback {
	return func(event *core.ChangeEvent, err *error) *error {
		if err != nil {
			return err
		}

		if event_types.Filter(event.Type, c.AllowedTypes) && core.ExtensionFilter(event.FileName, c.Extensions) {
			command := core.CommandTokens(c.Command, event)
			commandError := commandRunner(command)

			if commandError != nil && c.AllowFailing {
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
) *error {
	fmt.Printf("Watching path: %s and running command: '%s'\n", c.Paths, c.Command)
	callback := onFileChange(c, commandRunner)

	return fileWatcher(c.Paths, callback, context.Background())
}
