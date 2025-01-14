package utils

import (
	"io"
	"strings"
)

var _ io.Writer = &ObscuredWriter{}

type ObscuredWriter struct {
	w       io.Writer
	secrets []string
}

func NewObscuredWriter(wrappedWriter io.Writer, secrets []string) *ObscuredWriter {
	return &ObscuredWriter{
		w:       wrappedWriter,
		secrets: secrets,
	}
}

func (w *ObscuredWriter) Write(p []byte) (int, error) {
	original := p

	for _, secret := range w.secrets {
		if secret == "" {
			continue
		}

		p = []byte(strings.ReplaceAll(string(p), secret, "******"))
	}

	n, err := w.w.Write(p)
	if err != nil {
		return len(original), err
	}
	if n != len(p) {
		return len(original), io.ErrShortWrite
	}

	return len(original), nil
}
