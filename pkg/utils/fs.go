package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func IsSymlink(fi os.FileInfo) bool {
	return fi.Mode()&fs.ModeSymlink == fs.ModeSymlink
}

func IsPermModeMatched(stat fs.FileInfo, desiredPerm fs.FileMode) bool {
	return stat.Mode().Perm() == desiredPerm
}

func readSymbolicLinkUntilRealPath(path string) (string, error) {
	finalPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}

	return finalPath, nil
}

func ChmodIfUnmatched(logger *logrus.Entry, path string, stat fs.FileInfo, desiredPerm fs.FileMode) error {
	if IsPermModeMatched(stat, desiredPerm) {
		return nil
	}

	err := os.Chmod(path, desiredPerm)
	if err != nil {
		return fmt.Errorf("failed to chmod %s to %s: %w", path, desiredPerm, err)
	}

	return nil
}

func ChmodAndChownRecursively(logger *logrus.Entry, path string, uid int, gid int, mode os.FileMode) error {
	dir := path
	logger = logger.WithFields(logrus.Fields{"dir": path})

	stat, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		dir = filepath.Dir(path)
		logger = logger.WithField("dir", dir)
		logger.Warn("path is not a directory, fallback to parent directory instead")
	}

	err = ChmodIfUnmatched(logger, dir, stat, mode)
	if err != nil {
		return err
	}

	err = os.Chown(dir, uid, gid)
	if err != nil {
		return fmt.Errorf("failed to chown %s to %d:%d: %w", dir, uid, gid, err)
	}

	return fs.WalkDir(os.DirFS(dir), ".", func(walkPath string, walkDirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if walkPath == "." {
			return nil
		}
		if walkPath == ".." {
			return nil
		}

		walkPath = filepath.Join(dir, walkPath)

		stat, err := walkDirEntry.Info()
		if err != nil {
			return fmt.Errorf("failed to get info of %s: %w", walkPath, err)
		}

		if IsSymlink(stat) {
			resolvedWalkPath, err := readSymbolicLinkUntilRealPath(walkPath)
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}

				return fmt.Errorf("failed to resolve symlink %s: %w", walkPath, err)
			}

			stat, err = os.Stat(resolvedWalkPath)
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}

				return fmt.Errorf("failed to get info of resolved symlink %s: %w", resolvedWalkPath, err)
			}
		}
		// optionally chmod
		if stat.IsDir() || stat.Mode().IsRegular() {
			err = ChmodIfUnmatched(logger, walkPath, stat, mode)
			if err != nil {
				return err
			}
		}

		err = os.Chown(walkPath, uid, gid)
		if err != nil {
			return fmt.Errorf("failed to chown %s to %d:%d: %w", walkPath, uid, gid, err)
		}

		return nil
	})
}

func CleanupNotExistingSymlinks(logger *logrus.Entry, path string) error {
	return fs.WalkDir(os.DirFS(path), ".", func(walkPath string, walkDirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if walkPath == "." {
			return nil
		}
		if walkPath == ".." {
			return nil
		}

		walkPath = filepath.Join(path, walkPath)

		stat, err := walkDirEntry.Info()
		if err != nil {
			return fmt.Errorf("failed to get info of %s: %w", walkPath, err)
		}

		if IsSymlink(stat) {
			_, err := os.Stat(walkPath)
			if err != nil {
				if os.IsNotExist(err) {
					logger.Warnf("removing dangling symlink %s", walkPath)
					err = os.Remove(walkPath)
					if err != nil {
						return fmt.Errorf("failed to remove dangling symlink %s: %w", walkPath, err)
					}
				} else {
					return fmt.Errorf("failed to get info of symlink %s: %w", walkPath, err)
				}
			}
		}

		return nil
	})
}
