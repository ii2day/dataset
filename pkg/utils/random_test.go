package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt64(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		assert := assert.New(t)

		for i := 0; i < 1000; i++ {
			assert.NotZero(RandomInt64())
		}
	})

	t.Run("WithMax", func(t *testing.T) {
		assert := assert.New(t)

		for i := 0; i < 1000; i++ {
			assert.LessOrEqual(RandomInt64(10), int64(10))
		}
	})
}

func TestRandomBytes(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		assert := assert.New(t)

		rand1, _ := RandBytes()
		rand2, _ := RandBytes()

		assert.NotEqual(rand1, rand2)
		assert.Equal(len(rand1), 32)
		assert.Equal(len(rand2), 32)
	})
	t.Run("with args", func(t *testing.T) {
		assert := assert.New(t)

		arg := 123
		rand1, _ := RandBytes(arg)
		rand2, _ := RandBytes(arg)

		assert.NotEqual(rand1, rand2)
		assert.Equal(len(rand1), arg)
		assert.Equal(len(rand2), arg)
	})
}

func TestRandomHashString(t *testing.T) {
	assert := assert.New(t)

	hashString := RandomHashString()
	assert.NotEmpty(hashString)
	assert.Len(hashString, 64)

	hashString2 := RandomHashString(32)
	assert.NotEmpty(hashString2)
	assert.Len(hashString2, 32)
}
