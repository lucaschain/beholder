package infrastructure_test

import (
	"context"
	"os"
	"testing"

	"github.com/lucaschain/beholder/core"
	"github.com/lucaschain/beholder/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestFileWatcher(t *testing.T) {
	t.Run("should call callback when file is created", func(t *testing.T) {
		testFilePath := "/tmp/beholder"
		testFileName := "test.txt"
		testFile := testFilePath + "/" + testFileName
		ctx, cancel := context.WithCancel(context.Background())
		var fileName string
		callback := func(event *core.ChangeEvent, err *error) *error {
			fileName = event.FileName
			cancel()
			return nil
		}
		os.MkdirAll(testFilePath, 0755)
		os.Create(testFilePath + "/" + testFileName)
		defer os.RemoveAll(testFilePath)
		paths := []string{testFilePath}

		go func() {
			os.WriteFile(testFile, []byte("test"), 0644)
		}()

		infrastructure.FileWatcher(paths, callback, ctx)

		assert.Equal(t, testFile, fileName)
	})
}
