package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObscureString(t *testing.T) {
	str := ObscureString("test", []string{"test"})
	assert.Equal(t, "******", str)

	str = ObscureString("test-secret", []string{"secret"})
	assert.Equal(t, "test-******", str)
}
