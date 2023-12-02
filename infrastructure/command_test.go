package infrastructure_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/lucaschain/beholder/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	t.Run("runs a program capable of printing to a file", func(t *testing.T) {
		expected_string := "beholder is watching you"
		file_name := fmt.Sprintf("/tmp/%s_beholder_test.txt", fmt.Sprint(time.Now().Unix()))
		defer os.Remove(file_name)

		shell_command := fmt.Sprintf("echo -n %s > %s", expected_string, file_name)
		infrastructure.Command([]string{"sh", "-c", shell_command})

		file, err := os.ReadFile(file_name)
		if err != nil {
			t.Errorf("Error evaluating shell command file output: %s", err)
		}

		assert.Equal(t, expected_string, string(file))
	})
}
