package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitEngine(t *testing.T) {
	cfg := &Config{
		Output: "",
		Debug:  true,
	}
	InitEngine(cfg)
	assert.Equal(t, "debug", GetLevel().String())
}

func TestSetDebug(t *testing.T) {
	SetDebug()
	assert.Equal(t, "debug", GetLevel().String())
}
