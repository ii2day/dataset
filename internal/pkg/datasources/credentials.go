package datasources

import (
	"os"
	"path/filepath"

	"github.com/BaizeAI/dataset/pkg/log"
)

type Secrets struct {
	Username string `json:"-"`
	Password string `json:"-"`

	SSHPrivateKey           string `json:"-"`
	SSHPrivateKeyPassphrase string `json:"-"`

	Token string `json:"-"`

	AKSKAccessKeyID     string `json:"-"`
	AKSKSecretAccessKey string `json:"-"`
}

type SecretKey string

const (
	SecretKeyUsername             SecretKey = "username"
	SecretKeyPassword             SecretKey = "password"
	SecretKeyPrivateKey           SecretKey = "ssh-privatekey"
	SecretKeyPrivateKeyPassphrase SecretKey = "ssh-privatekey-passphrase" // #nosec G101
	SecretKeyToken                SecretKey = "token"
	SecretKeyAccessKey            SecretKey = "access-key"
	SecretKeySecretKey            SecretKey = "secret-key"
)

var (
	keys = []SecretKey{
		SecretKeyUsername,
		SecretKeyPassword,
		SecretKeyPrivateKey,
		SecretKeyPrivateKeyPassphrase,
		SecretKeyToken,
		SecretKeyAccessKey,
		SecretKeySecretKey,
	}
)

func ReadAndParseSecrets(name string) (Secrets, error) {
	mSecrets := make(map[SecretKey]string)

	logger := log.WithField("secretMountDir", name)

	for _, v := range keys {
		secretContent, err := os.ReadFile(filepath.Join(name, string(v)))
		if err != nil {
			logger.WithField("secretDataKey", v).Debug("failed to read secret")
			continue
		}

		mSecrets[v] = string(secretContent)
	}

	return Secrets{
		Username:                mSecrets[SecretKeyUsername],
		Password:                mSecrets[SecretKeyPassword],
		SSHPrivateKey:           mSecrets[SecretKeyPrivateKey],
		SSHPrivateKeyPassphrase: mSecrets[SecretKeyPrivateKeyPassphrase],
		Token:                   mSecrets[SecretKeyToken],
		AKSKAccessKeyID:         mSecrets[SecretKeyAccessKey],
		AKSKSecretAccessKey:     mSecrets[SecretKeySecretKey],
	}, nil
}
