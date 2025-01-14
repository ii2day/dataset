package datasources

import (
	"os"
)

type Options struct {
	// Primary arguments
	Type Type
	URI  string

	// --options flags
	Path string
	Mode os.FileMode
	UID  int
	GID  int
	Root string
}

type Loader interface {
	Sync(fromURI string, toPath string) error
}
