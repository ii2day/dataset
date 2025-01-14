// nolint: dupl
package datasources

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3Loader(t *testing.T) {
	loader, err := NewS3Loader(map[string]string{
		"region": "us-east-1",
	}, Options{
		Type: "",
		URI:  "s3://test-bucket",
		Path: "",
		Mode: 0,
		UID:  0,
		GID:  0,
		Root: "",
	}, Secrets{
		AKSKAccessKeyID:     "accid",
		AKSKSecretAccessKey: "acckey",
	})
	assert.NoError(t, err)
	fakeHTTP := fakeCommand{
		cmd: "rclone",
		outputs: []out{
			{
				stdout: "clone",
				stderr: "",
				exit:   0,
			},
			{
				stdout: "config",
				stderr: "",
				exit:   0,
			},
			{
				stdout: "config",
				stderr: "",
				exit:   0,
			},
		},
	}
	defer fakeHTTP.Clean()
	s3Dir, _ := os.MkdirTemp("", "s3Loader-*")
	defer os.RemoveAll(s3Dir)
	assert.NoError(t, err)
	fakeHTTP.WithContext(func() {
		err = loader.Sync("s3://test-bucket", s3Dir)
		assert.NoError(t, err)
	})
	bbs := fakeHTTP.GetAllInputs()
	assert.Equal(t, []byte("config touch\n"), bbs[0])
	assert.True(t, strings.HasPrefix(string(bbs[1]), "config create"))
	assert.True(t, strings.HasPrefix(string(bbs[2]), "sync"))
}
