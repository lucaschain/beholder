package event_types_test

import (
	"testing"

	"github.com/lucaschain/beholder/core/event_types"
	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	t.Run("should return event type from string", func(t *testing.T) {
		assert.Equal(t, event_types.Create, event_types.FromString("CREATE"))
		assert.Equal(t, event_types.Write, event_types.FromString("WRITE"))
		assert.Equal(t, event_types.Remove, event_types.FromString("REMOVE"))
		assert.Equal(t, event_types.Rename, event_types.FromString("RENAME"))
		assert.Equal(t, event_types.Chmod, event_types.FromString("CHMOD"))
	})
	t.Run("should return empty event type when string is invalid", func(t *testing.T) {
		assert.Equal(t, event_types.EventType(""), event_types.FromString("INVALID"))
	})
}
