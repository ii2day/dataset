package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestIsSymlink(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a regular file
	regularFile := filepath.Join(tempDir, "regular")
	err = os.WriteFile(regularFile, []byte("test"), 0600)
	assert.NoError(t, err)

	// Create a symlink
	symlinkFile := filepath.Join(tempDir, "symlink")
	err = os.Symlink(regularFile, symlinkFile)
	assert.NoError(t, err)

	// Test regular file
	regularInfo, err := os.Lstat(regularFile)
	assert.NoError(t, err)
	assert.False(t, IsSymlink(regularInfo))

	// Test symlink
	symlinkInfo, err := os.Lstat(symlinkFile)
	assert.NoError(t, err)
	assert.True(t, IsSymlink(symlinkInfo))
}

func TestIsPermModeMatched(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	err = tempFile.Chmod(0644)
	assert.NoError(t, err)

	info, err := os.Stat(tempFile.Name())
	assert.NoError(t, err)

	assert.True(t, IsPermModeMatched(info, 0644))
	assert.False(t, IsPermModeMatched(info, 0755))
}

func TestReadSymbolicLinkUntilRealPath(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Resolve the temporary directory to its real path
	realTempDir, err := filepath.EvalSymlinks(tempDir)
	assert.NoError(t, err)

	// Create a regular file
	regularFile := filepath.Join(realTempDir, "regular")
	err = os.WriteFile(regularFile, []byte("test"), 0600)
	assert.NoError(t, err)

	// Create a symlink to the regular file
	symlink1 := filepath.Join(realTempDir, "symlink1")
	err = os.Symlink(regularFile, symlink1)
	assert.NoError(t, err)

	// Create a symlink to the first symlink
	symlink2 := filepath.Join(realTempDir, "symlink2")
	err = os.Symlink(symlink1, symlink2)
	assert.NoError(t, err)

	// Test regular file
	path, err := readSymbolicLinkUntilRealPath(regularFile)
	assert.NoError(t, err)
	assert.Equal(t, regularFile, path)

	// Test symlink
	path, err = readSymbolicLinkUntilRealPath(symlink2)
	assert.NoError(t, err)
	assert.Equal(t, regularFile, path)

	// Test non-existent file
	_, err = readSymbolicLinkUntilRealPath(filepath.Join(realTempDir, "non-existent"))
	assert.Error(t, err)
}

func TestChmodIfUnmatched(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	logger := logrus.NewEntry(logrus.New())

	// Set initial permissions
	err = tempFile.Chmod(0644)
	assert.NoError(t, err)

	info, err := os.Stat(tempFile.Name())
	assert.NoError(t, err)

	// Test when permissions already match
	err = ChmodIfUnmatched(logger, tempFile.Name(), info, 0644)
	assert.NoError(t, err)

	// Test when permissions don't match
	err = ChmodIfUnmatched(logger, tempFile.Name(), info, 0755)
	assert.NoError(t, err)

	newInfo, err := os.Stat(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0755), newInfo.Mode().Perm())
}

func TestChmodAndChownRecursively(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	logger := logrus.NewEntry(logrus.New())

	err = os.Mkdir(filepath.Join(tempDir, "subdir"), 0755)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempDir, "file1"), []byte("test"), 0600)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(tempDir, "subdir", "file2"), []byte("test"), 0600)
	assert.NoError(t, err)

	err = ChmodAndChownRecursively(logger, tempDir, os.Getuid(), os.Getgid(), 0755)
	assert.NoError(t, err)

	// Check permissions (ownership can't be reliably tested without root)
	checkPerm := func(path string, expected os.FileMode) {
		info, err := os.Stat(path)
		assert.NoError(t, err)
		assert.Equal(t, expected, info.Mode().Perm())
	}

	checkPerm(tempDir, 0755)
	checkPerm(filepath.Join(tempDir, "subdir"), 0755)
	checkPerm(filepath.Join(tempDir, "file1"), 0755)
	checkPerm(filepath.Join(tempDir, "subdir", "file2"), 0755)
}

func TestCleanupNotExistingSymlinks(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	logger := logrus.NewEntry(logrus.New())

	// Create a regular file
	regularFile := filepath.Join(tempDir, "regular")
	err = os.WriteFile(regularFile, []byte("test"), 0600)
	assert.NoError(t, err)

	// Create a valid symlink
	validSymlink := filepath.Join(tempDir, "valid_symlink")
	err = os.Symlink(regularFile, validSymlink)
	assert.NoError(t, err)

	// Create a dangling symlink
	danglingSymlink := filepath.Join(tempDir, "dangling_symlink")
	err = os.Symlink(filepath.Join(tempDir, "non_existent"), danglingSymlink)
	assert.NoError(t, err)

	err = CleanupNotExistingSymlinks(logger, tempDir)
	assert.NoError(t, err)

	// Check that the valid symlink still exists
	_, err = os.Stat(validSymlink)
	assert.NoError(t, err)

	// Check that the dangling symlink has been removed
	_, err = os.Stat(danglingSymlink)
	assert.True(t, os.IsNotExist(err))
}
