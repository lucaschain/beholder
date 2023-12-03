package event_types_test

import (
	"testing"

	et "github.com/lucaschain/beholder/core/event_types"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("should return true only when allowed contains event type", func(t *testing.T) {
		allowed := []et.EventType{et.Create, et.Write}

		assert.True(t, et.Filter(et.Create, allowed))
		assert.True(t, et.Filter(et.Write, allowed))
		assert.False(t, et.Filter(et.Remove, allowed))
		assert.False(t, et.Filter(et.Rename, allowed))
		assert.False(t, et.Filter(et.Chmod, allowed))
	})
}
