package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	config *configuration
)

type configuration struct {
	DatasetJobSpecYaml string `json:"dataset_job_spec_yaml"`
}

func GetDatasetJobSpecYaml() string {
	if config == nil || config.DatasetJobSpecYaml == "" {
		return `
backoffLimit: 4
completionMode: NonIndexed
completions: 1
parallelism: 1
template:
  spec:
    restartPolicy: Never
    containers:
    - image: ubuntu:20.04
      command: ["/bin/bash", "-c", "echo 'Container args: '$(echo $@)"]
      #command: ["/bin/bash", "-c", "--"]
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 500m
          memory: 500Mi
`
	}
	return config.DatasetJobSpecYaml
}

func ParseConfigFromFileContent(content string) error {
	f, err := os.CreateTemp("", "dataset-config-*")
	if err != nil {
		return err
	}
	f.Write([]byte(content))
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	return ParseConfigFromFile(f.Name())
}

func ParseConfigFromFile(configPath string) error {
	cfg := &configuration{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	err := viper.Unmarshal(cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	})
	config = cfg
	if err != nil {
		return err
	}
	return nil
}
