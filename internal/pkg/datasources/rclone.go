package datasources

import (
	"os/exec"

	"github.com/BaizeAI/dataset/pkg/log"
	"github.com/BaizeAI/dataset/pkg/utils"
)

func rcloneCliConfigTouch() error {
	cmd := exec.Command("rclone", "config", "touch")
	logger := log.WithField("command", cmd.String())

	logger.Debug("executing command to touch rclone config")

	err := utils.ExecuteCommand(logger, cmd, nil)

	return err
}
