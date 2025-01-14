package conda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/BaizeAI/dataset/pkg/utils"
)

type MambaCLI struct {
	envs map[string]string
}

func NewMambaCLI() *MambaCLI {
	return &MambaCLI{
		envs: map[string]string{
			"always_yes": "true",
		},
	}
}

func (c *MambaCLI) newCommand(args ...string) *exec.Cmd {
	cmd := exec.Command("mamba", args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, c.GetEnvs()...)

	return cmd
}

func (c *MambaCLI) GetEnvs() []string {
	envs := make([]string, 0, len(c.envs))
	for k, v := range c.envs {
		envs = append(envs, fmt.Sprintf("%s=%s", fmt.Sprintf("MAMBA_%s", strings.ToUpper(k)), v))
		envs = append(envs, fmt.Sprintf("%s=%s", fmt.Sprintf("CONDA_%s", strings.ToUpper(k)), v))
	}

	return envs
}

// Version returns the version of conda
// Equivalent to `conda --version`
func (c *MambaCLI) Version(logger *logrus.Entry) (string, error) {
	args := []string{
		"--version",
	}

	cmd := c.newCommand(args...)
	output, err := utils.ExecuteCommandWithOutput(logger, cmd, []string{})
	if err != nil {
		return "", err
	}

	outputString := strings.TrimSpace(output.String())
	output.Reset()

	return strings.TrimSpace(outputString), nil
}

// Info returns the conda info
// Equivalent to `conda info --json`
func (c *MambaCLI) Info(logger *logrus.Entry) (*CondaInfoOutputRaw, error) {
	args := []string{
		"info",
		"--json",
	}

	cmd := c.newCommand(args...)
	output, err := utils.ExecuteCommandWithOutput(logger, cmd, []string{})
	if err != nil {
		return nil, err
	}

	defer output.Reset()

	var info CondaInfoOutputRaw
	err = json.Unmarshal(output.Bytes(), &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// EnvList returns the list of conda environments
// Equivalent to `conda env list`
func (c *MambaCLI) EnvList(logger *logrus.Entry) ([]CondaEnvListOutputEnv, error) {
	args := []string{
		"env",
		"list",
		"--json",
	}

	cmd := c.newCommand(args...)
	output, err := utils.ExecuteCommandWithOutput(logger, cmd, []string{})
	if err != nil {
		return make([]CondaEnvListOutputEnv, 0), err
	}

	defer output.Reset()

	var envList CondaEnvListOutputRaw
	err = json.Unmarshal(output.Bytes(), &envList)
	if err != nil {
		return make([]CondaEnvListOutputEnv, 0), err
	}

	envs := make([]CondaEnvListOutputEnv, 0, len(envList.Envs))
	for _, env := range envList.Envs {
		envs = append(envs, CondaEnvListOutputEnv{
			Name: filepath.Base(env),
			Path: env,
		})
	}

	return envs, nil
}

func (c *MambaCLI) ConfigSetShowChannelURLs(logger *logrus.Entry) {
	c.envs["show_channel_urls"] = "true"
}

// ConfigPrependPkgsDir prepends the pkgs_dir to the conda config
// Equivalent to `conda config --prepend pkgs_dirs <pkgsDir>`
func (c *MambaCLI) ConfigPrependPkgsDir(logger *logrus.Entry, pkgsDir string) {
	c.envs["pkgs_dirs"] = pkgsDir
}

// ConfigPrependEnvsDir prepends the envs_dir to the conda config
// Equivalent to `conda config --prepend envs_dirs <envsDir>`
func (c *MambaCLI) ConfigPrependEnvsDir(logger *logrus.Entry, envsDir string) {
	c.envs["envs_dirs"] = envsDir
}

var prefixAlreadyExistsErrRegexp = regexp.MustCompile(`^\sCondaValueError: prefix already exists: /.*\s\s$`)

func (c *MambaCLI) IsPrefixAlreadyExistsError(errBuffer *bytes.Buffer) bool {
	return prefixAlreadyExistsErrRegexp.Match(errBuffer.Bytes())
}

// CreateEnvFromFile creates a new conda environment from a file
// Equivalent to `conda env create --file <file> --verbose -y`
func (c *MambaCLI) CreateEnvFromFile(logger *logrus.Entry, file string) error {
	args := []string{
		"env",
		"create",
		"--file",
		file,
		"--verbose",
	}

	cmd := c.newCommand(args...)
	_, errBuffer, err := utils.ExecuteCommandWithAllOutput(logger, cmd, []string{})
	if err != nil {
		if c.IsPrefixAlreadyExistsError(errBuffer) {
			return nil
		}

		return err
	}

	return nil
}

// CleanAll cleans all conda packages
// Equivalent to `conda clean --all -y`
func (c *MambaCLI) CleanAll(logger *logrus.Entry) error {
	args := []string{
		"clean",
		"--all",
		"-y",
	}

	cmd := c.newCommand(args...)
	err := utils.ExecuteCommand(logger, cmd, []string{})
	if err != nil {
		return err
	}

	return nil
}
