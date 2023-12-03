package use_case_test

import (
	"testing"

	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/core/event_types"
	"github.com/lucaschain/beholder/use_case"
	"github.com/stretchr/testify/assert"
)

func TestWatch(t *testing.T) {
	t.Run("should only call callback if type is allowed", func(t *testing.T) {
		config := buildWatchConfig()
		config.AllowedTypes = []string{string(event_types.Create)}

		var called bool
		fakeRunner := func(command []string) error {
			called = true
			return nil
		}
		fakeWatcher := buildFakeFileWatcher("test.txt", event_types.Create)

		use_case.Watch(config, fakeWatcher, fakeRunner)

		assert.True(t, called)
	})

	t.Run("should not call callback if type is not allowed", func(t *testing.T) {
		config := buildWatchConfig()
		config.AllowedTypes = []string{string(event_types.Chmod)}

		var called bool
		fakeRunner := func(command []string) error {
			called = true
			return nil
		}
		fakeWatcher := buildFakeFileWatcher("test.txt", event_types.Create)

		use_case.Watch(config, fakeWatcher, fakeRunner)

		assert.False(t, called)
	})

	t.Run("should replace the command tokens based on the event", func(t *testing.T) {
		config := buildWatchConfig()
		config.AllowedTypes = []string{string(event_types.Create)}
		config.Command = []string{"echo", "{type}", "{file}"}

		var command []string
		fakeRunner := func(cmd []string) error {
			command = cmd
			return nil
		}
		fakeWatcher := buildFakeFileWatcher("test.txt", event_types.Create)

		use_case.Watch(config, fakeWatcher, fakeRunner)

		assert.Equal(t, []string{"echo", "CREATE", "test.txt"}, command)
	})
}

func buildFakeFileWatcher(filename string, eventType event_types.EventType) use_case.FileWatcher {
	if filename == "" {
		filename = "test.txt"
	}

	event := &core.ChangeEvent{
		Type:     event_types.Create,
		FileName: filename,
	}
	return func(paths []string, callback core.ChangeCallback) {
		callback(event, nil)
	}
}

func buildFakeRunner() use_case.CommandRunner {
	return func(command []string) error {
		return nil
	}
}

func buildWatchConfig() use_case.WatchConfig {
	return use_case.WatchConfig{
		Paths:        []string{"."},
		Command:      []string{"echo", "hello"},
		AllowedTypes: []string{event_types.Create.String()},
		AllowFailing: false,
	}
}
