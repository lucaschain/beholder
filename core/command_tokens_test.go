package core_test

import (
	"testing"

	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/core/event_types"
	"github.com/stretchr/testify/assert"
)

func TestCommandTokens(t *testing.T) {
	t.Run("should replace all strings in command", func(t *testing.T) {
		list := []string{
			"/foo/command",
			"--path",
			"{file}",
			"--type",
			"{type}",
		}
		event := core.ChangeEvent{
			FileName: "file.txt",
			Type:     event_types.Write,
		}

		tokens := core.CommandTokens(list, &event)

		expected := []string{
			"/foo/command",
			"--path",
			"file.txt",
			"--type",
			"WRITE",
		}

		assert.ElementsMatch(t, expected, tokens)
	})
}
