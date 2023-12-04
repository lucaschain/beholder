package core_test

import (
	"testing"

	"github.com/lucaschain/beholder/core"
	"github.com/stretchr/testify/assert"
)

func TestExtensionFilter(t *testing.T) {
	t.Run("should return true if no extensions passed", func(t *testing.T) {
		assert.True(t, core.ExtensionFilter("foo.go", []string{}))
	})

	t.Run("should return true if path has one of the extensions", func(t *testing.T) {
		assert.True(t, core.ExtensionFilter("foo.go", []string{".go", ".rb"}))
	})

	t.Run("should return false if the patch matches none of the extensions", func(t *testing.T) {
		assert.False(t, core.ExtensionFilter("foo.go", []string{".rb", ".js"}))
	})
}
