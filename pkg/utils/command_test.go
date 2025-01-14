package utils

import (
	"os"
	"os/exec"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/BaizeAI/dataset/pkg/log"
)

func TestExecuteCommandWithAllOutput(t *testing.T) {
	logger := log.WithFields(logrus.Fields{
		"test": "TestExecuteCommandWithAllOutput",
	})
	d, _ := os.MkdirTemp("", "fakeCommand-*")
	defer os.RemoveAll(d)
	os.WriteFile(d+"/test_output_0", []byte("output"), 0600)
	os.WriteFile(d+"/test_output_1", []byte("error"), 0600)
	t.Run("run with secret", func(t *testing.T) {
		o, _, err := ExecuteCommandWithAllOutput(logger, exec.Command("ls", d), []string{"test_output_1"})
		assert.NoError(t, err)
		assert.Equal(t, "test_output_0\n******\n", o.String())
	})
}
