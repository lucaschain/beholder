package event_types_test

import (
	"testing"

	"github.com/lucaschain/beholder/core/event_types"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("should return true only when allowed contains event type", func(t *testing.T) {
		allowed := []string{"CREATE", "WRITE"}
		assert.True(t, event_types.Filter(event_types.Create, allowed))
		assert.True(t, event_types.Filter(event_types.Write, allowed))
		assert.False(t, event_types.Filter(event_types.Remove, allowed))
		assert.False(t, event_types.Filter(event_types.Rename, allowed))
		assert.False(t, event_types.Filter(event_types.Chmod, allowed))
	})
}
