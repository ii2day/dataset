package utils

import (
	"strings"
)

func ObscureString(str string, secrets []string) string {
	for _, v := range secrets {
		if (v) == "" {
			continue
		}

		str = strings.ReplaceAll(str, v, "******")
	}

	return str
}
